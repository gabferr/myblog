package handlers

import (
	"net/http"
	"text/template"

	"github.com/gabferr/myblog/db"
)

// Exibe a página inicial com os posts mais recentes
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}

	// Obtém a lista de posts
	posts, err := db.GetAllPosts(db.DBConn)
	if err != nil {
		http.Error(w, "Erro ao carregar posts", http.StatusInternalServerError)
		return
	}

	// Renderiza a página com os posts
	tmpl.Execute(w, posts)
}
