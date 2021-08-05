package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fernandocao/ideazmxgoweb/internal/config"
	"github.com/fernandocao/ideazmxgoweb/internal/models"
	"github.com/fernandocao/ideazmxgoweb/internal/render"
)

//Repo es usado por los handlers
var Repo *Repository

//Repository el el tipo Repository
type Repository struct {
	App *config.AppConfig
}

//NewRopo crea un nuevo repositorio
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers establce el repositorio para los handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `jason:"ok"`
	Message string `json:"message"`
}

func (m *Repository) PostHome(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Recibido",
	}

	out, err := json.MarshalIndent(resp, "", "      ")

	if err != nil {
		log.Print(err)
	}
	/*
		name := r.Form.Get("name")
		email := r.Form.Get("email")
		subject := r.Form.Get("subject")
		message := r.Form.Get("message")
	*/
	//w.Write([]byte(fmt.Sprintf("%s envío del correo %s, asunto: %s, mensaje: %s", name, email, subject, message)))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

/*
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
*/
