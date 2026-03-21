package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Produto struct{
	ID int `json: "id"`
	Nome string `json: "nome"`
	Preco float64 `json: "preco"`
}

var produtos = []Produto{
	{ID: 1, Nome: "Mouse", Preco: 50.0},
	{ID: 2, Nome: "Teclado", Preco: 150.0},

}
func listarProdutos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(produtos)
}
func main(){
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "pong")
	})

	http.HandleFunc("/produtos",listarProdutos)

	fmt.Println("Servidor rodando em :8080")
	http.ListenAndServe(":8080",nil)
}