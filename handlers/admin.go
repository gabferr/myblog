package handlers

import (
	"net/http"
	"text/template"

	"github.com/gabferr/myblog/db"
)

// AdminHandler exibe o painel de administração com a lista de posts
func AdminHandler(w http.ResponseWriter, r *http.Request) {
    // Busca todos os posts do banco de dados
    posts, err := db.GetAllPosts(dbConn)
    if err != nil {
        http.Error(w, "Erro ao carregar posts", http.StatusInternalServerError)
        return
    }

    // Carrega o template da página de administração
    tmpl, err := template.ParseFiles("templates/admin.html")
    if err != nil {
        http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
        return
    }

    // Renderiza o template, passando os posts como dados
    err = tmpl.Execute(w, posts)
    if err != nil {
        http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
        return
    }
}
