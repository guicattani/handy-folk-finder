package helpers

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/guicattani/handy-folk-finder/internal/config"
)

var ac *config.AppConfig

// NewHelpers sets up app config for helpers
func NewHelpers(a *config.AppConfig) {
	ac = a
}

// ClientError shows errors with no logging, ideal for 4xx errors
func ClientError(w http.ResponseWriter, status int, msg string) {
	ac.InfoLog.Println("Client error with status of", status)
	ac.InfoLog.Println("Error", msg)
	http.Error(w, http.StatusText(status), status)
}

// ServerError shows errors with logging, ideal for 5xx errors
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ac.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// DbConfig returns DB url considering the .env file
func DbConfig() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"))
}
