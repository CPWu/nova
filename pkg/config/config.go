package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application configuration settings, including a template cache that maps template names to their corresponding parsed templates. This allows for efficient rendering of templates by avoiding the need to parse them on each request.
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
}
