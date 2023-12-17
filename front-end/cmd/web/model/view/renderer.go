package view

import (
	"fmt"
	"net/http"
	"html/template"
)

func RenderTemplate(w http.ResponseWriter, t string, data map[string]interface{}) {

	partials := []string{
		fmt.Sprintf("./cmd/web/templates/%s", t),
		"./cmd/web/templates/base.layout.gohtml",
		"./cmd/web/templates/header.partial.gohtml",
		"./cmd/web/templates/footer.partial.gohtml",
		"./cmd/web/templates/my-account.partial.gohtml",
	}

	tmpl, err := template.ParseFiles(partials...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}