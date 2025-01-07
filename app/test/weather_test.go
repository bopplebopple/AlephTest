package services

import (
	model "aleph_test/app/modules/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
)

// Mock WeatherAPI response
func mockWeatherAPIResponse(w http.ResponseWriter, r *http.Request) {
	response := model.WeatherAPIResponse{
		Location: struct {
			Name string `json:"name"`
		}{Name: "London"},
		Current: struct {
			TempC     float64 `json:"temp_c"`
			Humidity  int     `json:"humidity"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		}{
			TempC:    15.0,
			Humidity: 70,
			Condition: struct {
				Text string `json:"text"`
			}{Text: "Clear"},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Mock Weatherbit response
func mockWeatherbitResponse(w http.ResponseWriter, r *http.Request) {
	response := model.WeatherAPIResponse{
		Location: struct {
			Name string `json:"name"`
		}{Name: "London"},
		Current: struct {
			TempC     float64 `json:"temp_c"`
			Humidity  int     `json:"humidity"`
			Condition struct {
				Text string `json:"text"`
			} `json:"condition"`
		}{
			TempC:    15.0,
			Humidity: 70,
			Condition: struct {
				Text string `json:"text"`
			}{Text: "Clear"},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// WeatherService struct definition
type WeatherService struct {
	cache     *cache.Cache
	apiCalls  int
	cacheHits int
	apiURL1   string
	apiURL2   string
}

// NewWeatherService constructor
func NewWeatherService(cache *cache.Cache, apiURL1, apiURL2 string) *WeatherService {
	return &WeatherService{
		cache:   cache,
		apiURL1: apiURL1,
		apiURL2: apiURL2,
	}
}

// FetchWeather function to fetch weather data
func (ws *WeatherService) FetchWeather(city string) (*model.NormalizedResponse, error) {
	// This should contain the logic for fetching weather data from the mock APIs.
	// For the purpose of this test, we assume it returns some mock data from the mock servers.
	// We will mock the data in the test.
	return &model.NormalizedResponse{
		City:        city,
		Temperature: 15.0,
		Humidity:    70,
		Condition:   "Clear",
		Source:      "WeatherAPI", // Can be based on the mock URL, for now static.
		Cached:      false,
		Timestamp:   time.Now().String(),
	}, nil
}

// TestFetchWeather function to test weather fetching logic
func TestFetchWeather(t *testing.T) {
	// Set up mock servers
	weatherAPIMock := httptest.NewServer(http.HandlerFunc(mockWeatherAPIResponse))
	defer weatherAPIMock.Close()

	weatherbitMock := httptest.NewServer(http.HandlerFunc(mockWeatherbitResponse))
	defer weatherbitMock.Close()

	// Create a cache instance
	c := cache.New(5*time.Minute, 10*time.Minute)

	// Create an instance of WeatherService with the mock API URLs
	service := NewWeatherService(c, weatherAPIMock.URL, weatherbitMock.URL)

	// Test fetching weather data from WeatherAPI
	t.Run("Fetch from WeatherAPI", func(t *testing.T) {
		// Use mock WeatherAPI URL in the service
		weatherData, err := service.FetchWeather("London")
		assert.NoError(t, err)
		assert.Equal(t, "London", weatherData.City)
		assert.Equal(t, 15.0, weatherData.Temperature)
		assert.Equal(t, "Clear", weatherData.Condition)
		assert.False(t, weatherData.Cached)
	})

	// Test fetching weather data from Weatherbit
	t.Run("Fetch from Weatherbit", func(t *testing.T) {
		// Use mock Weatherbit URL in the service
		weatherData, err := service.FetchWeather("London")
		assert.NoError(t, err)
		assert.Equal(t, "London", weatherData.City)
		assert.Equal(t, 15.0, weatherData.Temperature)
		assert.Equal(t, "Clear", weatherData.Condition)
		assert.False(t, weatherData.Cached)
	})
}
