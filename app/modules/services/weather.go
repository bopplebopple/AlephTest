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
	apiKey    string
}

func NewWeatherService(cache *cache.Cache, apiKey string) *WeatherService {
	return &WeatherService{
		cache:    cache,
		apiKey:   apiKey,
		apiCalls: 0,
	}
}

func (s *WeatherService) FetchWeather(city string) (*model.NormalizedResponse, error) {
	if cachedData, found := s.cache.Get(city); found {
		return cachedData.(*model.NormalizedResponse), nil
	}

	s.apiCalls++

	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", s.apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %d", resp.StatusCode)
	}

	var apiResponse model.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode weather data: %v", err)
	}

	normalized := &model.NormalizedResponse{
		City:        apiResponse.Location.Name,
		Temperature: apiResponse.Current.TempC,
		Humidity:    apiResponse.Current.Humidity,
		Condition:   apiResponse.Current.Condition.Text,
		Source:      "WeatherAPI",
		Cached:      false,
		Timestamp:   time.Now().UTC().Format(time.RFC3339),
	}

	cacheDuration := time.Duration(30 * time.Minute)
	s.cache.Set(city, normalized, cacheDuration)

	return normalized, nil
}

func (s *WeatherService) GetStats() *model.Stats {
	return &model.Stats{
		APICalls:  s.apiCalls,
		CacheHits: s.cacheHits,
	}
}
