package models

type Configuration struct {
	LogFile                string `json:"logFile"`
	ClientLogFile          string `json:"clientLogFile"`
	IsProduction           string `json:"isProduction"`
	ApiKey                 string `json:"apiKey"`
	ConnectionString       string `json:"connectionString"`
	GraphAPIPort           string `json:"graphAPIPort"`
	GrpcAPIPort            string `json:"gRPCAPIPort"`
	DatabaseName           string `json:"databaseName"`
	SecureModeGRPC         string `json:"secureModeGRPC"`
	SSLCertificateLocation string `json:"sslCertificateLocation"`
	SSLKeyLocation         string `json:"sslKeyLocation"`
}
