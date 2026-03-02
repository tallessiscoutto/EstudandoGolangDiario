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

	pessoaMaisVelha := Pessoas[0]
	for _, Pessoa := range Pessoas {
		if Pessoa.Idade > pessoaMaisVelha.Idade {
			pessoaMaisVelha = Pessoa
			
		}	
	}
	fmt.Println("Pessoa mais velha ",pessoaMaisVelha)

}