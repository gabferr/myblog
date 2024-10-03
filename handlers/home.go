package handlers

import (
	"log"
	"net/http"

	"github.com/gabferr/myblog/db"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts(db.DBConn)
	if err != nil {
		log.Printf("Erro ao obter posts: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	data := struct {
		Posts []db.Post
	}{
		Posts: posts,
	}

	renderTemplate(w, "home", data)
}
