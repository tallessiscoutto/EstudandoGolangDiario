package main

import "fmt"

 var normal int
 var tipo_ataque *string
type Personagem struct {
	Nome string
	Vida int
	Tipo_Dano string
	Tipo_ataque string
	Dano int
}
type Monstro struct {
	Tipo string
	Vida int
}


func(p *Personagem) Apresentar() {
	fmt.Println("Meu nome é",p.Nome, "e tenho",p.Vida,"de vida")
}
func(m *Monstro)Aparecimento(){
	fmt.Println("Olha apareceu um",m.Tipo,"mate-o")
}
func condicaoataque(p *Personagem,m *Monstro) int{
	if (p.Tipo_ataque != m.Tipo){
		fmt.Println("Ataque Super Efetivo!")
		return p.Dano * 2
	} 
	fmt.Println("Ataque Normal")
	return p.Dano
}
func ReceberDano(m *Monstro, dano int){
	m.Vida -= dano	
}

func main(){

	p := Personagem{Nome: "Talles", Vida:100,Tipo_ataque: "Morte",Dano: 15};
	monstro := Monstro{Tipo: "Sangue", Vida:50};
	p.Apresentar()
	monstro.Aparecimento()
	for monstro.Vida > 0 {
		fmt.Println(p.Nome, "Ataca")
		dano := condicaoataque(&p, &monstro)
		ReceberDano(&monstro, dano)
		fmt.Println("O monstro recebeu",dano,"de dano")
		fmt.Println("Vida restante",monstro.Vida)
	}
	fmt.Println("Monstro Derrotado!")
}



