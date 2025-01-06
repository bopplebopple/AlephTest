package main

import (
	"aleph_test/app/middlewares"
	"aleph_test/app/modules/handlers"
	"aleph_test/app/modules/routes"
	"aleph_test/app/modules/services"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/patrickmn/go-cache"
)

func main() {
	middlewares.InitRateLimiter()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cacheExpiryTimeStr := os.Getenv("CACHE_EXPIRY_TIME")
	cacheExpiryTime, err := strconv.Atoi(cacheExpiryTimeStr)
	if err != nil {
		log.Fatal("Invalid CACHE_EXPIRY_TIME environment variable")
	}

	cacheExpiryDuration := time.Duration(cacheExpiryTime) * time.Minute

	c := cache.New(cacheExpiryDuration, cacheExpiryDuration)

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY environment variable is required")
	}

	weatherService := services.NewWeatherService(c, apiKey)

	weatherHandler := handlers.NewWeatherHandler(weatherService)

	r := gin.Default()

	routes.SetupRoutes(r, weatherHandler)

	r.Run(":8080")
}
