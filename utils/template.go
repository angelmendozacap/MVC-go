package utils

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("./views/*.html"))
}

// ExecuteTemplate escutes the html templates
func ExecuteTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl.ExecuteTemplate(w, name, data)
}
