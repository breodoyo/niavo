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
	JWTSecret string
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET environment variable is required")
	}

	return Config{
		AppName:     appName,
		Env:         env,
		Port:        ":" + port,
		DatabaseURL: databaseUrl,
		JWTSecret: jwtSecret,
	}
}
