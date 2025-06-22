package handlers

import (
	"net/http"

	"github.com/Aakash-1857/HotelBooking/pkg/config"
	"github.com/Aakash-1857/HotelBooking/pkg/models"
	"github.com/Aakash-1857/HotelBooking/pkg/render"
)

var Repo *Repository
type Repository struct{
	App *config.AppConfig
}
func NewRepo (a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r* Repository){
	Repo = r
}
// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	stringMap := make(map[string]string)
	stringMap["test"] = "HEY SEXY!"
	render.RenderTemplate(w, "home.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	newMap := make(map[string]string)
	newMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl",&models.TemplateData{
		StringMap: newMap,
	})
}
