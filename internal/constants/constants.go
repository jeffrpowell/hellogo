package constants

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

// Registered environment vars
const (
	ENV_PORT              string = "PORT"
	ENV_POSTGRES_USER     string = "POSTGRES_USER"
	ENV_POSTGRES_PASSWORD string = "POSTGRES_PASSWORD"
	ENV_POSTGRES_HOST     string = "POSTGRES_HOST"
	ENV_POSTGRES_DB       string = "POSTGRES_DB"
)

// Handler consts
const defaultPort string = "8080"

var PORT string = loadEnvWithDefault(ENV_PORT, defaultPort)
var ROUTER *mux.Router = mux.NewRouter()

func loadEnvWithDefault(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		val = defaultValue
	}
	return val
}

// Database consts
const (
	DB_DEFAULT_USER        string = "hellogo"
	DB_DEFAULT_PASSWORD    string = "hellogo"
	DB_DEFAULT_HOST        string = "localhost"
	DB_DEFAULT_DB          string = "hellogo"
	DB_TABLE_COLORGRADIENT string = "hellogo.colorgradient"
)

var DB_CONNECTION_STRING string = getDbConnectionString()

func getDbConnectionString() string {
	// Fetch database connection parameters from environment variables
	dbUser := loadEnvWithDefault(ENV_POSTGRES_USER, DB_DEFAULT_USER)
	dbPassword := loadEnvWithDefault(ENV_POSTGRES_PASSWORD, DB_DEFAULT_PASSWORD)
	dbHost := loadEnvWithDefault(ENV_POSTGRES_HOST, DB_DEFAULT_HOST)
	dbName := loadEnvWithDefault(ENV_POSTGRES_DB, DB_DEFAULT_DB)

	// Construct the connection string
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbName)
}
