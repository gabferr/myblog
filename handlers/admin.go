package handlers

import (
	"net/http"

	"github.com/gabferr/myblog/db"
)

// AdminHandler exibe o painel de administração com a lista de posts
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	// Busca todos os posts do banco de dados usando a conexão exportada
	posts, err := db.GetAllPosts(db.DBConn)
	if err != nil {
		http.Error(w, "Erro ao carregar posts", http.StatusInternalServerError)
		return
	}

	// Prepara os dados para passar ao template
	data := map[string]interface{}{
		"Posts": posts,
	}

	// Usa a função renderTemplate para renderizar a página admin
	renderTemplate(w, "admin", data)
}
