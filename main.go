package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabferr/myblog/db"
	"github.com/gabferr/myblog/handlers"
)

func main() {
    // Inicializa a conexão com o banco de dados
    dbConn, err := db.Initialize()
    if err != nil {
        log.Fatalf("Erro ao inicializar o banco de dados: %v", err)
    }
    defer dbConn.Close()

    // Passando o banco de dados para os handlers
    handlers.SetDB(dbConn)

    // Configuração das rotas
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/post", handlers.PostHandler)
    http.HandleFunc("/admin", handlers.AdminHandler)
    http.HandleFunc("/login", handlers.LoginHandler)
    http.HandleFunc("/register", handlers.RegisterHandler)
    http.HandleFunc("/create", handlers.CreatePostHandler)

    // Inicializa o servidor
    fmt.Println("Servidor rodando em http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
