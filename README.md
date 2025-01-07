# 🌦️ Weather Service
Hello Aleph-Labs,

My name is Matthew, and I want to thank you for giving me the opportunity to create this service for you. Even though I didn't receive full marks for the bonus point, I would still like to express my gratitude. This test has taught me a lot and will provide me with valuable things to study in the future.

The following is a Go-based weather service that fetches real-time weather data for a given city from multiple weather APIs. The service leverages caching to reduce API calls and improve performance. It supports fetching weather data from two different weather API providers: **WeatherAPI** and **Weatherbit**.

## 🚀 Features

- 🌍 Fetches weather data for a city by querying **WeatherAPI** and **Weatherbit** APIs.
- 🕒 Caches the results for 30 minutes to reduce API calls and improve response time.
- 📊 Provides API call statistics, including the total number of API calls and cache hits.
- ⚠️ Handles errors gracefully and attempts to fetch data from multiple sources when one source fails.

## 💻 Installation

### 📦 Prerequisites

- 🐋 Docker
- 🐳 Docker Compose (optional for multi-container setups)

To run this project, ensure you have [Docker](https://www.docker.com/) installed on your machine.

### 🛠️ Installation

1. **Install Docker**  
   Follow the installation instructions on the [official Docker website](https://www.docker.com/).

2. **Create a `.env` file**  
   Create a `.env` file in the root of the project directory and include the following environment variables:
   ```dotenv
   WEATHER_BIT_API_KEY=your-weather-bit-api-key
   WEATHER_API_KEY=your-weather-api-key
   RATE_LIMIT_PER_SECOND=1
   RATE_LIMIT_BURST_SIZE=20
   CACHE_EXPIRY_TIME=30

3. **Build the Docker Image**  
   Open your terminal and run the following command to build the Docker image:

   ```bash
   docker build -t bopple/aleph_test .
   ```
   
4. **Run the Docker Container**  
   After building the image, you can run the Docker container with:

   ```
   docker run -d --name myapp -p 3000:3000 bopple/aleph_test
   ```

## 🌐 Access the Application
Once the Docker container is running, you can access the application at http://localhost:3000.
