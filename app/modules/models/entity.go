package model

type WeatherAPIResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity  int     `json:"humidity"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type NormalizedResponse struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
	Condition   string  `json:"condition"`
	Source      string  `json:"source"`
	Cached      bool    `json:"cached"`
	Timestamp   string  `json:"timestamp"`
}

type Stats struct {
	APICalls  int `json:"api_calls"`
	CacheHits int `json:"cache_hits"`
}
