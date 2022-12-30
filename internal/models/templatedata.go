package models

import "github.com/urschmidt/bookings/internal/forms"

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	Form      *forms.Form
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
