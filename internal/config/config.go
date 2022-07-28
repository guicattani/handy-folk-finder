package config

import (
	"log"
	"text/template"
)

// AppConfig holds the application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
	InProduction  bool
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
}
