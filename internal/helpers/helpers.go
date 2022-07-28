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

func ClientError(w http.ResponseWriter, status int) {
	ac.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	ac.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func DbConfig() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("DEVELOPMENT_DB_HOST"),
		os.Getenv("DEVELOPMENT_DB_PORT"),
		os.Getenv("DEVELOPMENT_DB"),
		os.Getenv("DEVELOPMENT_DB_USERNAME"),
		os.Getenv("DEVELOPMENT_DB_PASSWORD"))
}
