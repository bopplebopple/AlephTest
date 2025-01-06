package routes

import (
	"aleph_test/app/modules/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, weatherHandler *handlers.WeatherHandler) {

	r.GET("/weather", weatherHandler.GetWeather)
	r.GET("/stats", weatherHandler.GetStats)
}
