package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabferr/myblog/db"
	"github.com/gabferr/myblog/handlers"
	"github.com/gabferr/myblog/middleware"
)

func main() {
	// Inicializa a conexão com o banco de dados
	dbConn, err := db.Initialize()
	if err != nil {
		log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
	}
	defer dbConn.Close()

	// Passando o banco de dados para os handlers
	db.SetDB(dbConn)

	// Configuração das rotas
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/admin", middleware.BasicAuth(handlers.AdminHandler, "admin", "admin"))
	http.HandleFunc("/login", middleware.BasicAuth(handlers.LoginHandler, "admin", "admin"))
	http.HandleFunc("/register", middleware.BasicAuth(handlers.RegisterHandler, "admin", "admin"))
	http.HandleFunc("/admin/create-post", middleware.BasicAuth(handlers.CreatePostHandler, "admin", "admin"))

	// Inicializa o servidor
	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
