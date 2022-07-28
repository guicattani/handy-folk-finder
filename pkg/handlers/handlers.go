package handlers

import (
	"net/http"

	"github.com/guicattani/go-course/pkg/config"
	"github.com/guicattani/go-course/pkg/models"
	"github.com/guicattani/go-course/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there!"

	render.RenderTemplate(w, "home.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
