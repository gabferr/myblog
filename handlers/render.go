package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

/*
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    // Define the layout and content templates
    layoutPath := filepath.Join("templates", "layout.html")
    contentPath := filepath.Join("templates", tmpl+".html")

    // Parse both templates
    t, err := template.ParseFiles(layoutPath, contentPath)
    if err != nil {
        log.Printf("Erro ao analisar templates: %v", err)
        http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
        return
    }

    // Execute the template, using the layout as the base
    err = t.ExecuteTemplate(w, "layout", data)
    if err != nil {
        log.Printf("Erro ao executar template: %v", err)
        http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
        return
    }*/

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Define the layout, content, and navbar templates
	layoutPath := filepath.Join("templates", "layout.html")
	navbarPath := filepath.Join("templates", "navbar.html")
	contentPath := filepath.Join("templates", tmpl+".html")

	// Parse all the necessary templates (layout, navbar, and content)
	t, err := template.ParseFiles(layoutPath, navbarPath, contentPath)
	if err != nil {
		log.Printf("Erro ao analisar templates: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Execute the layout template, which will include the navbar and content
	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Erro ao executar template: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}
}
