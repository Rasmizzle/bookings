package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

type TemplateCache map[string]*template.Template

// AppConfig holds the application wide configuration
type AppConfig struct {
	UseCache        bool
	TemplateCache   TemplateCache
	InfoLog         *log.Logger
	DevelopmentMode bool
	Session         *scs.SessionManager
}
