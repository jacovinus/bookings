// TemplateData holds data sent from handlers to templates
package models

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRSFToken string
	Flash     string
	Warn      string
	Error     string
}
