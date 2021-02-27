package handlers

import (
	"github.com/senny-matrix/bookings/pkg/config"
	"github.com/senny-matrix/bookings/pkg/models"
	"github.com/senny-matrix/bookings/pkg/render"
	"net/http"
)



// Repo the repository used by the handlers
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

// Home is a home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl.html", &models.TemplateData{})
}

// About is the about page handler.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// Perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	// Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl.html", &models.TemplateData{StringMap: stringMap})
}

// addValue adds two integers and returns the sum
func AddValue(x, y int) int {
	return x + y
}
