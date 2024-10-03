package handlers

import (
	"log"
	"net/http"

	"github.com/gabferr/myblog/db"
	"github.com/gabferr/myblog/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetAllPosts(db.DBConn)
	if err != nil {
		log.Printf("Erro ao obter posts: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Converte a fatia de ponteiros em uma fatia de valores
	var postValues []models.Post
	for _, post := range posts {
		postValues = append(postValues, *post) // Desreferencia o ponteiro para obter o valor
	}

	data := struct {
		Posts []models.Post
	}{
		Posts: postValues,
	}

	renderTemplate(w, "home", data)
}
