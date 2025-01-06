package middlewares

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

// RateLimiter is the rate limiter instance
var RateLimiter *rate.Limiter

// Load environment variables and initialize the rate limiter
func InitRateLimiter() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetch rate limit settings from environment variables
	rateLimitPerSecond, err := strconv.Atoi(os.Getenv("RATE_LIMIT_PER_SECOND"))
	if err != nil {
		log.Fatal("Error parsing RATE_LIMIT_PER_SECOND")
	}

	burstSize, err := strconv.Atoi(os.Getenv("RATE_LIMIT_BURST_SIZE"))
	if err != nil {
		log.Fatal("Error parsing RATE_LIMIT_BURST_SIZE")
	}

	// Create a new rate limiter with the loaded values
	RateLimiter = rate.NewLimiter(rate.Every(time.Second/time.Duration(rateLimitPerSecond)), burstSize)
}

// RateLimitMiddleware applies rate limiting to requests
func RateLimitMiddleware(c *gin.Context) {
	if !RateLimiter.Allow() {
		c.JSON(429, gin.H{"error": "Rate limit exceeded, please try again later"})
		c.Abort() // Stop the request here
		return
	}
	c.Next() // Proceed with the next handler if the rate limit is not exceeded
}
