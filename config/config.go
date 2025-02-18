package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Host           string
	Port           string
	DatabaseURI    string
	AllowedOrigins []string

	PgUser     string
	PgPassword string
	DbName     string
}

// Generate a new config for use in the server, using environment variables
// already on the system, or using default values for local development
func New() Config {
	host := getEnvDefault("HOST", "localhost")
	port := getEnvDefault("PORT", "7777")

	databaseURI := getEnvDefault("DATABASE_URI", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	ao := strings.Split(getEnvDefault("ALLOWED_ORIGINS", fmt.Sprintf("http://%s:%s,https://%s:%s", host, port, host, port)), ",")

	pgUser := getEnvDefault("POSTGRES_USERNAME", "postgres")
	pgPass := getEnvDefault("POSTGRES_PASSWORD", "postgres")
	pgDb := getEnvDefault("POSTGRES_DATABASE", "postgres")

	return Config{
		Host:           host,
		Port:           port,
		DatabaseURI:    databaseURI,
		AllowedOrigins: ao,
		PgUser:         pgUser,
		PgPassword:     pgPass,
		DbName:         pgDb,
	}
}

// this function allows easy switching between a local dev
// and a cloud provider environment variables
func getEnvDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}
