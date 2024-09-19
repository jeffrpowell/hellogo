package constants

import (
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
