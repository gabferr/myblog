package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gabferr/myblog/db"
)

// Exibe a página de um post individual
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Extrai o ID do post da query string (ex: /post?id=123)
	postID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil || postID < 1 {
		http.Error(w, "ID de post inválido", http.StatusBadRequest)
		return
	}

	// Busca o post pelo ID no banco de dados
	post, err := db.GetPostByID(db.DBConn, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post não encontrado", http.StatusNotFound)
		} else {
			http.Error(w, "Erro ao carregar post", http.StatusInternalServerError)
		}
		return
	}
	renderTemplate(w, "post", post)

}
