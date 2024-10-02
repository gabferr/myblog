package middleware

import (
	"net/http"
)

func BasicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		// Verifica se a autenticação foi fornecida e se as credenciais estão corretas
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Se as credenciais estiverem corretas, chama o próximo handler
		next.ServeHTTP(w, r)
	}
}
