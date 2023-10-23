package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //The CSRF token is included in forms or as a header in HTTP requests. For example, it might be included as a hidden input field in an HTML form, or in an HTTP header.
	Flash     string
	Warning   string
	Error     string
}