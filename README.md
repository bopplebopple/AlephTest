# 🌦️ Weather Service

A Go-based weather service that fetches real-time weather data for a given city from multiple weather APIs. The service leverages caching to reduce API calls and improve performance. It supports fetching weather data from two different weather API providers: **WeatherAPI** and **Weatherbit**.

## 🚀 Features

- 🌍 Fetches weather data for a city by querying **WeatherAPI** and **Weatherbit** APIs.
- 🕒 Caches the results for 30 minutes to reduce API calls and improve response time.
- 📊 Provides API call statistics, including the total number of API calls and cache hits.
- ⚠️ Handles errors gracefully and attempts to fetch data from multiple sources when one source fails.

## 💻 Installation

### 📦 Prerequisites

- 🐋 Docker
- 🐳 Docker Compose (optional for multi-container setups)

### 🛠️ Steps