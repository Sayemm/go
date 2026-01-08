package config

import (
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("SERVICE_NAME", "test-service")
	os.Setenv("HTTP_PORT", "8080")
	os.Setenv("JWT_SECRET_KEY", "test-secret")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpass")

	cnf := GetConfig()

	if cnf.ServiceName != "test-service" {
		t.Errorf("Expected service name 'test-service', got %s", cnf.ServiceName)
	}

	if cnf.HttpPort != 8080 {
		t.Errorf("Expected port 8080, got %d", cnf.HttpPort)
	}

	if cnf.DB.Host != "localhost" {
		t.Errorf("Expected DB host 'localhost', got %s", cnf.DB.Host)
	}
}

// go test ./config
