package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMode bool
}

type Config struct {
	Version      string
	ServiceName  string
	Environment  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() error {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using system environment variables")
	}

	version := getEnvOrDefault("VERSION", "1.0.0")
	serviceName := getEnvRequired("SERVICE_NAME")
	environment := getEnvOrDefault("ENVIRONMENT", "development")
	httpPort := getEnvAsIntRequired("HTTP_PORT")
	jwtSecretKey := getEnvRequired("JWT_SECRET_KEY")

	dbConfig := &DBConfig{
		Host:          getEnvRequired("DB_HOST"),
		Port:          getEnvAsIntRequired("DB_PORT"),
		Name:          getEnvRequired("DB_NAME"),
		User:          getEnvRequired("DB_USER"),
		Password:      getEnvRequired("DB_PASSWORD"),
		EnableSSLMode: getEnvAsBool("DB_ENABLE_SSL_MODE", false),
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		Environment:  environment,
		HttpPort:     httpPort,
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

	return nil
}

func GetConfig() *Config {
	if configurations == nil {
		if err := loadConfig(); err != nil {
			fmt.Println("Failed to load configurations: ", err)
			os.Exit(1)
		}
	}
	return configurations
}

func getEnvRequired(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Error: Required environment variable %s is not set\n", key)
		os.Exit(1)
	}
	return value
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsIntRequired(key string) int {
	valueStr := getEnvRequired(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error: Environment variable %s must be a number, got: %s\n", key, valueStr)
		os.Exit(1)
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		fmt.Printf("Warning: Invalid boolean for %s, using default: %v\n", key, defaultValue)
		return defaultValue
	}
	return value
}
