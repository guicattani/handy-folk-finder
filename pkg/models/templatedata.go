package models

// TemplateData holds data for templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float     map[string]float32
	Data      map[string]interface{}
	CSFRToken string
	Flash     string
	Warning   string
	Error     string
}
