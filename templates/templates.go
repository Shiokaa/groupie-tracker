package templates

import (
	"fmt"
	"html/template"
)

var Temp *template.Template

func InitTemplates() {
	parsedTemplate, err := template.ParseGlob("templates/*.html")

	if err != nil {
		fmt.Println("Error parsing templates", err)

	}

	Temp = parsedTemplate
}
