package middlewares

import (
	"fmt"
	"log"
	"net/http"
)

// logger escreve informacoes da requisicao no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// autenticar verifica se o usuario fazendo a requisicao esta autenticado
func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("autenticando....")
		next(w, r)

	}
}
