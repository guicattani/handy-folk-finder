package handlers

import (
	"net/http"

	"github.com/guicattani/handy-folk-finder/internal/config"
	"github.com/guicattani/handy-folk-finder/internal/driver"
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
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there!"
}
