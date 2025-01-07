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

To run this project, ensure you have [Docker](https://www.docker.com/) installed on your machine.

### 🛠️ Installation

1. **Install Docker**  
   Follow the installation instructions on the [official Docker website](https://www.docker.com/).

2. **Build the Docker Image**  
   Open your terminal and run the following command to build the Docker image:

   ```bash
   docker build -t bopple/aleph_test .
   ```
   
3. **Run the Docker Container**  
   After building the image, you can run the Docker container with:

   ```
   docker run -d --name myapp -p 3000:3000 bopple/aleph_test
   ```

## 🌐 Access the Application
Once the Docker container is running, you can access the application at http://localhost:3000.