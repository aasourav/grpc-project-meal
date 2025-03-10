package config

import "os"

type Config struct {
	ServerPort string
	DBURI      string
}

func LoadConfig() *Config {
	port := "8008"
	dbURI := "mongodb://localhost:27017/meal"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("DBURI") != "" {
		dbURI = os.Getenv("DBURI")
	}

	return &Config{
		ServerPort: port,
		DBURI:      dbURI,
	}
}
