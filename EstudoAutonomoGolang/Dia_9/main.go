package main

import "fmt"


type Personagem struct {
	Nome string
	Vida int
}

func(p Personagem) Apresentar() {
	fmt.Println("Meu nome é",p.Nome, "e tenho",p.Vida,"de vida")
}

func ReceberDano(p Personagem, dano int)Personagem {
	p.Vida -= dano
	return p
}
func(p *Personagem) Curar(cura int){
	p.Vida += cura
}
func main(){
	p1 := Personagem{Nome: "Talles", Vida:100}

	//Método
	p1.Apresentar()
	//Função
	p1 = ReceberDano(p1, 30)

	p1.Apresentar()

	p1.Curar(20)

	p1.Apresentar()
}



