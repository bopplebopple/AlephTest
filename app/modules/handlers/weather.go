package handlers

import (
	service "aleph_test/app/modules/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	service *service.WeatherService
}

func NewWeatherHandler(service *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetWeather(c *gin.Context) {
	city := c.DefaultQuery("city", "Jakarta")
	weatherData, err := h.service.FetchWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, weatherData)
}

func (h *WeatherHandler) GetStats(c *gin.Context) {
	stats := h.service.GetStats()
	c.JSON(http.StatusOK, stats)
}
