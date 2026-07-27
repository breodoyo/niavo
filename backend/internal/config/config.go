package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Port        string
	Env         string
	DatabaseURL string
}

// Load returns the application configuration
func Load() Config {
	_ = godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "Niavi_API"
	}
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	databaseUrl := os.Getenv("DATABASE_URL")

	return Config{
		AppName:     appName,
		Env:         env,
		Port:        ":" + port,
		DatabaseURL: databaseUrl,
	}
}
