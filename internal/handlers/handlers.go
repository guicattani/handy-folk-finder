package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/guicattani/handy-folk-finder/internal/config"
	"github.com/guicattani/handy-folk-finder/internal/driver"
	"github.com/guicattani/handy-folk-finder/internal/helpers"
	"github.com/guicattani/handy-folk-finder/internal/repository"
	"github.com/guicattani/handy-folk-finder/internal/repository/dbrepo"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// NewRepo creates a new repository
func NewRepo(ac *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: ac,
		DB:  dbrepo.NewPostgresRepo(db.Connection, ac),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	partners, err := m.DB.AllPartners()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	out, err := json.MarshalIndent(partners, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
