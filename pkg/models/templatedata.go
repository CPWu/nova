package models

// TemplateData struct that holds various types of data that can be passed to templates when rendering them. It includes maps for strings, integers, and floats, as well as a generic map for any other type of data. Additionally, it includes fields for CSRF tokens, flash messages, error messages, and warning messages that can be used to provide feedback to the user when rendering templates.
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float64
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Error     string
	Warning   string
}
