package main

import "fmt"

type Jogador struct {
	Nome  string
	Gols int
}

func main() {

	Jogadores := []Jogador{
		{
			Nome:  "Talles",
			Gols: 1,
		},
		{
			Nome:  "Neymar",
			Gols: 403,
		},
		{
			Nome:  "Messi",
			Gols: 893,
		},
	}
	jogadoresCadastrados(Jogadores)
	artilheiro(Jogadores)
	totalGols(Jogadores)
}
func jogadoresCadastrados(Jogadores []Jogador){
fmt.Println("Jogadores")
	for i := 0; i < len(Jogadores); i++ {
		fmt.Println(Jogadores[i].Nome)
	}
}
func artilheiro(Jogadores []Jogador){
	var jogadorArtilheiro Jogador
	for _, Jogador := range Jogadores{
		if Jogador.Gols > jogadorArtilheiro.Gols {
			jogadorArtilheiro = Jogador
		}
	}
	fmt.Println("Jogador com mais Gols", jogadorArtilheiro)
}
func totalGols(Jogadores []Jogador){
	var somaGols = 0
	for _, Jogador := range Jogadores{
		somaGols += Jogador.Gols 
	}
	fmt.Println("Total de gols de todos os jogadores ", somaGols)
}