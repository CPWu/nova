package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/cpwu/nova/pkg/config"
	"github.com/cpwu/nova/pkg/models"
)

// functions holds the custom template functions that can be used in the templates. It is currently empty, but can be populated with any necessary functions for template rendering.
var functions = template.FuncMap{}

// app is a package-level variable that holds the application configuration, including the template cache. It is used to access the template cache when rendering templates.
var app *config.AppConfig

// NewTemmplates creates a new template cache and assigns it to the AppConfig struct. This allows the application to efficiently render templates by storing parsed templates in memory, avoiding the need to parse them on each request.
func NewTemmplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData is a helper function that can be used to add default data to the TemplateData struct before rendering a template. This allows for consistent data to be available in all templates without having to manually add it each time.
func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var tc map[string]*template.Template

	// Get the template cache from the app config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("template not found in cache:", ok)
	}

	// Create a buffer to hold the rendered template
	buf := new(bytes.Buffer)
	templateData = AddDefaultData(templateData)
	err := t.Execute(buf, templateData)
	if err != nil {
		log.Println("error executing template:", err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("error writing template to response:", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// Get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Loop through the pages found and create a template for each one
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
