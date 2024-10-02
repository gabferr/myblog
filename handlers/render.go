package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Cache de templates compilados
// Compilando o template globalmente para evitar recompilação em cada request
var templates = template.Must(template.ParseFiles(
	"templates/layout.html", // O layout deve ser carregado primeiro
	"templates/home.html",
	"templates/admin.html",
	"templates/login.html",
	"templates/post.html",
	"templates/register.html",
))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data) // Use o nome do template que você deseja renderizar
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		log.Println("Erro ao renderizar template:", err)
	}
}
