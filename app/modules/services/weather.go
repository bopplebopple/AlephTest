package services

import (
	model "aleph_test/app/modules/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type WeatherService struct {
	cache     *cache.Cache
	apiCalls  int
	cacheHits int
	apiKey1   string
	apiKey2   string
}

func NewWeatherService(cache *cache.Cache, apiKey1, apiKey2 string) *WeatherService {
	return &WeatherService{
		cache:    cache,
		apiKey1:  apiKey1,
		apiKey2:  apiKey2,
		apiCalls: 0,
	}
}

func (s *WeatherService) FetchWeather(city string) (*model.NormalizedResponse, error) {
	if cachedData, found := s.cache.Get(city); found {
		s.cacheHits++
		cachedData.(*model.NormalizedResponse).Cached = true
		return cachedData.(*model.NormalizedResponse), nil
	}

	s.apiCalls++

	responseChan := make(chan *model.NormalizedResponse, 2)
	errorChan := make(chan error, 2)

	coordinates, err := getCityCoordinates(city, s.apiKey1)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch city coordinates: %v", err)
	}

	go func() {
		url1 := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.apiKey1, city)
		resp, err := http.Get(url1)
		if err != nil {
			errorChan <- fmt.Errorf("WeatherAPI error: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errorChan <- fmt.Errorf("WeatherAPI returned status: %d", resp.StatusCode)
			return
		}

		var apiResponse model.WeatherAPIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			errorChan <- fmt.Errorf("WeatherAPI decode error: %v", err)
			return
		}

		responseChan <- &model.NormalizedResponse{
			City:        apiResponse.Location.Name,
			Temperature: apiResponse.Current.TempC,
			Humidity:    apiResponse.Current.Humidity,
			Condition:   apiResponse.Current.Condition.Text,
			Source:      "WeatherAPI",
			Cached:      false,
			Timestamp:   time.Now().UTC().Format(time.RFC3339),
		}
	}()

	go func() {
		url2 := fmt.Sprintf(
			"https://api.weatherbit.io/v2.0/current?lat=%f&lon=%f&key=%s&units=metric",
			coordinates.Latitude, coordinates.Longitude, s.apiKey2,
		)
		resp, err := http.Get(url2)
		if err != nil {
			errorChan <- fmt.Errorf("Weatherbit error: %v", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			errorChan <- fmt.Errorf("Weatherbit returned status: %d", resp.StatusCode)
			return
		}

		var apiResponse struct {
			Data []struct {
				Temp    float64 `json:"temp"`
				Rh      int     `json:"rh"`
				Weather struct {
					Description string `json:"description"`
				} `json:"weather"`
			} `json:"data"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
			errorChan <- fmt.Errorf("Weatherbit decode error: %v", err)
			return
		}

		if len(apiResponse.Data) == 0 {
			errorChan <- fmt.Errorf("No weather data found")
			return
		}

		responseChan <- &model.NormalizedResponse{
			City:        city,
			Temperature: apiResponse.Data[0].Temp,
			Humidity:    apiResponse.Data[0].Rh,
			Condition:   apiResponse.Data[0].Weather.Description,
			Source:      "Weatherbit",
			Cached:      false,
			Timestamp:   time.Now().UTC().Format(time.RFC3339),
		}
	}()

	select {
	case res := <-responseChan:
		s.cache.Set(city, res, 30*time.Minute)
		return res, nil
	case err := <-errorChan:
		select {
		case res := <-responseChan:
			s.cache.Set(city, res, 30*time.Minute)
			return res, nil
		case err2 := <-errorChan:
			return nil, fmt.Errorf("both APIs failed: WeatherAPI: %v, OpenWeather: %v", err, err2)
		case <-time.After(5 * time.Second):
			return nil, fmt.Errorf("second API call timed out after first failure: %v", err)
		}
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("API calls timed out")
	}
}

func (s *WeatherService) GetStats() *model.Stats {
	return &model.Stats{
		APICalls:  s.apiCalls,
		CacheHits: s.cacheHits,
	}
}

func getCityCoordinates(city, apiKey string) (*model.Coordinates, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/search.json?key=%s&q=%s", apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch coordinates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	var result []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode coordinates: %v", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no coordinates found for city: %s", city)
	}

	return &model.Coordinates{
		Latitude:  result[0].Lat,
		Longitude: result[0].Lon,
	}, nil
}
