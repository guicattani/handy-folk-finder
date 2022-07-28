package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// NewHandlers creates new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// AllPartners renders all partners from the DB as a JSON
func (m *Repository) AllPartners(w http.ResponseWriter, r *http.Request) {
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

// AllPartners renders a specific partner from the DB as a JSON
func (m *Repository) SpecificPartner(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		helpers.ClientError(w, http.StatusUnprocessableEntity, "id parameter missing")
		return
	}

	id, err := strconv.Atoi(keys[0])
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	partner, err := m.DB.SpecificPartner(id)
	if err != nil {
		helpers.ClientError(w, http.StatusNotFound, "partner not found")
		return
	}

	out, err := json.MarshalIndent(partner, "", "     ")
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// AllPartners renders a all partner close to a given user as a JSON
func (m *Repository) ClosestPartner(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()

	if len(keys) < 3 || keys["email"] == nil || keys["password"] == nil || keys["needed_experience"] == nil {
		helpers.ClientError(w, http.StatusUnprocessableEntity, "missing parameters, must have email, login, needed_experience")
		return
	}

	customer, err := m.DB.CustomerLogin(keys["email"][0], keys["password"][0])
	if err != nil {
		helpers.ClientError(w, http.StatusNotFound, "customer not found")
		return
	}

	partners, err := m.DB.ClosestPartner(customer, keys["needed_experience"][0])
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
