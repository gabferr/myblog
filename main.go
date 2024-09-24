package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Post struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
}

var posts = []Post{
	{ID: 1, Title: "Meu primeiro post", Content: "Este é o conteúdo do primeiro post.", CreatedAt: time.Now()},
}

var nextID = 2

const (
	adminUsername = "admin"
	adminPassword = "senha123" // Na prática, use uma senha forte e armazene de forma segura
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/create", createPostHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := r.URL.Query().Get("id")
	for _, post := range posts {
		if fmt.Sprintf("%d", post.ID) == id {
			err = tmpl.Execute(w, post)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.Error(w, "Post não encontrado", http.StatusNotFound)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUsername && password == adminPassword {
			http.SetCookie(w, &http.Cookie{
				Name:    "authenticated",
				Value:   "true",
				Expires: time.Now().Add(24 * time.Hour),
			})
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("authenticated")
	if err != nil || cookie.Value != "true" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, posts)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("authenticated")
	if err != nil || cookie.Value != "true" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")

		newPost := Post{
			ID:        nextID,
			Title:     title,
			Content:   content,
			CreatedAt: time.Now(),
		}
		posts = append(posts, newPost)
		nextID++

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}
