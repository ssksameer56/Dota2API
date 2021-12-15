package models

type Configuration struct {
	LogFile      string `json:"logFile"`
	IsProduction bool   `json:"isProduction"`
	ApiKey       string `json:"apiKey"`
}
