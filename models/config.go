package models

type Configuration struct {
	LogFile          string `json:"logFile"`
	IsProduction     string `json:"isProduction"`
	ApiKey           string `json:"apiKey"`
	ConnectionString string `json:"connectionString"`
	GraphAPIPort     string `json:"graphAPIPort"`
	GrpcAPIPort      string `json:"gRPCAPIPort"`
	DatabaseName     string `json:"databaseName"`
}
