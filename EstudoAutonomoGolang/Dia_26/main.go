package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Produto struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Preco float64 `json:"preco"`
}

var produtos = []Produto{
	{ID: 1, Nome: "Mouse", Preco: 50.0},
	{ID: 2, Nome: "Teclado", Preco: 150.0},
}

func listarProdutos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(produtos)
}
func criarProduto(w http.ResponseWriter, r *http.Request) {
	var novo Produto

	//transforma json em struct
	err := json.NewDecoder(r.Body).Decode(&novo)
	if err != nil {
		http.Error(w, "Erro ao ler dados", http.StatusBadRequest)
		return
	}
	produtos = append(produtos, novo)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novo)
}
func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "pong",
		})
	})

	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			listarProdutos(w, r)
		} else if r.Method == "POST" {
			criarProduto(w, r)
		}
	})
	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080", nil)
}
