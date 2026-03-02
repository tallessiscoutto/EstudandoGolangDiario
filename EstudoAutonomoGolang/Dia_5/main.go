package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {

	Pessoas := []Pessoa{
		{
			Nome:  "Talles",
			Idade: 21,
		},
		{
			Nome:  "Julio",
			Idade: 15,
		},
		{
			Nome:  "Renato",
			Idade: 25,
		},
	}
	fmt.Println("Pessoas")
	for i := 0; i < len(Pessoas); i++ {
		fmt.Println(Pessoas[i].Nome)
	}
	fmt.Println("Pessoas maiores de idade")
	for i, _ := range Pessoas{
		if Pessoas[i].Idade >= 18 {
			fmt.Println(Pessoas[i])
		}
	}

	pessoaMaisVelha := Pessoas[0]
	for _, Pessoa := range Pessoas {
		if Pessoa.Idade > pessoaMaisVelha.Idade {
			pessoaMaisVelha = Pessoa
			
		}	
	}
	fmt.Println("Pessoa mais velha ",pessoaMaisVelha)
	calcularIdadeMedia(Pessoas)
}
func calcularIdadeMedia(pessoas []Pessoa) {
	var soma int 
	soma = 0
	for _, p := range pessoas {
		soma += p.Idade
	}
	var media float64 = float64(soma) / float64(len(pessoas))
	fmt.Println("Media de Idades : ", media)
}