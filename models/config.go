package models

type Configuration struct {
	LogFile          string `json:"logFile"`
	IsProduction     bool   `json:"isProduction"`
	ApiKey           string `json:"apiKey"`
	ConnectionString string `json:"connString"`
	GraphAPIPort     string `json:"graphAPIPort"`
	DatabaseName     string `json:"databaseName"`
}
