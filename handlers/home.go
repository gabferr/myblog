package handlers

import (
	"log"
	"net/http"

	"github.com/gabferr/myblog/db"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts()
	if err != nil {
		http.Error(w, "Erro ao obter posts", http.StatusInternalServerError)
		log.Println("Erro ao obter posts:", err)
		return
	}

	// Renderiza o template 'home.html' dentro do layout
	renderTemplate(w, "home.html", posts)
}
