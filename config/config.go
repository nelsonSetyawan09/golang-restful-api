package config

import "os"

type Config struct {
	AppHost string
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	APIKey            string
	CorsAllowedOrigin string
}

func LoadConfig() *Config {
	return &Config{
		AppHost: os.Getenv("APP_HOST"),
		AppPort: os.Getenv("APP_PORT"),

		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),

		APIKey:            os.Getenv("API_KEY"),
		CorsAllowedOrigin: os.Getenv("CORS_ALLOWED_ORIGIN"),
	}
}
