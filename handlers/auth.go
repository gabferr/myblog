package handlers

import (
	"net/http"

	"github.com/gabferr/myblog/db"
	"github.com/gabferr/myblog/models"
)

// LoginHandler exibe a página de login e processa o formulário de login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(w, "login", nil)
	case http.MethodPost:
		handleLoginPost(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func handleLoginPost(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Validação e autenticação do usuário
	user, err := db.GetUserByUsername(db.DBConn, username)
	if err != nil {
		http.Error(w, "Erro ao processar login", http.StatusInternalServerError)
		return
	}

	if user == nil || user.Password != password {
		// Renderiza a página de login novamente com uma mensagem de erro
		data := map[string]interface{}{
			"Error": "Usuário ou senha incorretos",
		}
		renderTemplate(w, "login", data)
		return
	}

	// TODO: Implementar uma forma segura de gerenciar sessões
	// Por exemplo, usando github.com/gorilla/sessions

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

// Exibe a página de registro e processa o formulário de registro
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "register", nil)

	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		// Verifica se o usuário já existe
		user, _ := db.GetUserByUsername(db.DBConn, username)
		if user != nil {
			http.Error(w, "Nome de usuário já em uso", http.StatusBadRequest)
			return
		}

		// Cria o usuário no banco de dados
		newUser := &models.User{
			Username: username,
			Password: password,
			Email:    email,
		}
		err := db.CreateUser(db.DBConn, newUser)
		if err != nil {
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}

		// Redireciona para a página de login após registro
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// Cria um novo post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		renderTemplate(w, "create_post", nil)
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		// TODO: Validar os dados do formulário

		err := db.CreatePost(db.DBConn, &models.Post{Title: title, Content: content})
		if err != nil {
			renderTemplate(w, "create-post", map[string]interface{}{
				"Error": "Erro ao criar o post. Por favor, tente novamente.",
			})
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
