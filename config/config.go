package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config = Config{}

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}
	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}

/*
main.go - Process Create
Bind the Env variable to the process
Process is like a computer
	- if we write env to a computer it shows all the env file
*/
