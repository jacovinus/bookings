package handlers

import (
	"net/http"

	"github.com/jacovinus/bookings/pkg/config"
	"github.com/jacovinus/bookings/pkg/models"
	"github.com/jacovinus/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplateOne(w, "home.page.tmpl", &models.TemplateData{})

}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello gain"
	// send some data to the template

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remoteIP"] = remoteIP
	render.RenderTemplateOne(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
