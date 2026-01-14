package db

import (
	"ecommerce/config"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s",
		cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Name,
	)
	if !cnf.EnableSSLMode {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	connString := GetConnectionString(cnf)

	fmt.Println("Connecting to database...")
	fmt.Printf("Host: %s:%d\n", cnf.Host, cnf.Port)
	fmt.Printf("Database: %s\n", cnf.Name)

	maxRetries := 10
	retryDelay := 2 * time.Second

	var dbCon *sqlx.DB
	var err error

	for i := range maxRetries {
		dbCon, err = sqlx.Connect("postgres", connString)
		if err == nil {
			err = dbCon.Ping()
			if err == nil {
				fmt.Println("Database Connected Successfully!")
				return dbCon, nil
			}
		}

		if i < maxRetries-1 {
			fmt.Printf("Connection attempt %d/%d failed:%v\n", i+1, maxRetries, err)
			fmt.Printf("Retrying in %v...seconds\n", retryDelay)
			time.Sleep(retryDelay)
		}
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}
