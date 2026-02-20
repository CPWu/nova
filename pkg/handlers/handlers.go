package handlers

import (
	"net/http"

	"github.com/cpwu/nova/pkg/config"
	"github.com/cpwu/nova/pkg/models"
	"github.com/cpwu/nova/pkg/render"
)

// Repo is the repository used by the handlers. It holds a reference to the application configuration, allowing the handlers to access the template cache and other settings when rendering templates and handling requests.
var Repo *Repository

// Repository struct that holds a reference to the application configuration. This allows the handlers to access the template cache and other configuration settings when rendering templates and handling requests.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository with the given application configuration. This function is used to initialize the repository that will be used by the handlers to access the application configuration and template cache.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers. This function is called in the main function to initialize the handlers with the repository that contains the application configuration and template cache.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler that writes a simple response to the client and handles any potential errors that may occur during the writing process. It also prints the number of bytes written to the console for debugging purposes.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About page handler that demonstrates the use of a helper function to add two values and print the result. It also writes a response to the client and handles any potential errors that may occur during the writing process.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
