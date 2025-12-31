package config

import (
	"os"
	"strconv"
)

type Environment string

const (
	EnvDevelopment Environment = "development"
	EnvStaging     Environment = "staging"
	EnvProduction  Environment = "production"
	EnvTest        Environment = "test"
)

type Config struct {
	AppEnv       Environment
	Port         string
	DatabaseURL  string
	JWTSecret    string
	Debug        bool
	AllowOrigins string
}

func LoadConfig() *Config {
	return &Config{
		AppEnv:       getEnvironment(),
		Port:         getEnv("PORT", "8080"),
		DatabaseURL:  getEnv("DATABASE_URL", ""),
		JWTSecret:    getEnv("JWT_SECRET", "dev-secret"),
		Debug:        getEnvBool("DEBUG", false),
		AllowOrigins: getEnv("ALLOW_ORIGINS", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return parsed
}

func getEnvironment() Environment {
	env := Environment(os.Getenv("APP_ENV"))
	switch env {
	case EnvDevelopment, EnvStaging, EnvProduction, EnvTest:
		return env
	default:
		return EnvDevelopment
	}
}
