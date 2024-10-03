package handlers

import (
    "html/template"
    "log"
    "net/http"
    "path/filepath"
)

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
    }
}