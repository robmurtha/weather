# OpenWeather Service

# API Calls

## Hourly Weather
The hourly weather data can be retrieved by latitude and longitude.
https://your-server.com/weather/current/hourly?lat=44.34&lon=10.99

## Hourly Weather API data types

### GetCurrentWeatherRequest
Internal struct for holding query parameters.
```
type GetCurrentWeatherRequest struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}
```

### GetCurrentWeatherResponse
Result type for returning as JSON.
```
type GetCurrentWeatherResponse struct {
	CurrentConditions *open_weather.CurrentConditions `json:"currentConditions,omitempty"`
	Alerts            []string                        `json:"alerts,omitempty"`

	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
	Error      error  `json:"error,omitempty"`
}
```
### CurrentConditions
```                           
{
"conditions" {
  "coord": {
    "lon": 10.99,
    "lat": 44.34
  },
  "weather": [
    {
      "id": 501,
      "main": "Rain",
      "description": "moderate rain",
      "icon": "10d"
    }
  ],
  "base": "stations",
  "main": {
    "temp": 298.48,
    "feels_like": 298.74,
    "temp_min": 297.56,
    "temp_max": 300.05,
    "pressure": 1015,
    "humidity": 64,
    "sea_level": 1015,
    "grnd_level": 933
  },
  "visibility": 10000,
  "wind": {
    "speed": 0.62,
    "deg": 349,
    "gust": 1.18
  },
  "rain": {
    "1h": 3.16
  },
  "clouds": {
    "all": 100
  },
  "dt": 1661870592,
  "sys": {
    "type": 2,
    "id": 2075663,
    "country": "IT",
    "sunrise": 1661834187,
    "sunset": 1661882248
  },
  "timezone": 7200,
  "id": 3163858,
  "name": "Zocca",
  "cod": 200
},
"error"
}                      
```

# TODO
* unit tests
* curl tests
