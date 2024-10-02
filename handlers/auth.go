package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"fmt"

	"github.com/gabferr/myblog/db"
	"github.com/gabferr/myblog/models"
)

// Exibe a página de login e processa o formulário de login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Imprime o username e password recebidos (para debug)
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		// Validação e autenticação do usuário
		user, err := db.GetUserByUsername(db.DBConn, username)
		if err != nil || user == nil || user.Password != password {
			http.Error(w, "Usuário ou senha incorretos", http.StatusUnauthorized)
			return
		}

		// Definir sessão ou cookies para autenticar o usuário
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}

// Exibe a página de registro e processa o formulário de registro
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
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
		tmpl, err := template.ParseFiles("templates/create_post.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")
		userID, _ := strconv.ParseInt(r.FormValue("user_id"), 10, 64) // Exemplo: pegar o ID do usuário logado

		// Cria um post
		newPost := &models.Post{UserID: userID, Title: title, Content: content}

		err := db.CreatePost(db.DBConn, newPost)
		if err != nil {
			http.Error(w, "Erro ao criar post", http.StatusInternalServerError)
			return
		}

		// Redireciona para a página inicial ou de administração
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
