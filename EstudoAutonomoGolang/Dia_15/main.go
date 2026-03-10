package main

import (
	"fmt"
	"math/rand"
)
	
var direcao string
var aldeao string
type Personagem struct{
	Nome string;
	Vida int
	Tipo_ataque string
	Dano int
	Andar int
}
type Monstro struct{
	Nome string
	Tipo string
	Vida int}
func (p *Personagem) Apresentar(){
	fmt.Println("Olá eu sou o ",p.Nome,"Seja bem vindo a essa jornada guerreiro, me ajude a derrotar os inimigos que aparecerem pelo caminho")
}
func EscolherDirecao( p *Personagem){
	fmt.Println("Escolha uma direção 1 = Esquerda e 2 = Direita")
	fmt.Scanln(&p.Andar)
}
func Aparecer(monstros[]Monstro)Monstro{
	indice := rand.Intn(len(monstros))
	m := monstros[indice]
	fmt.Println("Cuidado, esse caminho é meio sombrio")
	fmt.Println("Arbustos se mexem ")
	fmt.Println("Meu deus um ",m.Nome,"é hora do combate guerreiro")
	return m
}
func vilareijo(p* Personagem, aldeao string){
	fmt.Println("Olá caro viajante, meu nome é",aldeao,"Se estiver a procura de missões eu tenho uma boa pra você")
	aldeao = "Tadeu"
}

func verificarCritico(p *Personagem, m *Monstro)int{
	if(p.Tipo_ataque != m.Tipo){
		fmt.Println("Critico")
		return p.Dano *2
	}
	fmt.Println("dano comum")
	return p.Dano
}
type Alvo interface{
	ReceberDano(dano int)
	
}

func (p *Personagem)Atacar(a Alvo, dano int){
	a.ReceberDano(p.Dano)
}
func (m *Monstro) ReceberDano(dano int){
	m.Vida -= dano
}
func main(){
	p1 := Personagem{Nome: "Talles",Vida:100,Tipo_ataque:"Sangue",Dano:15,}
	monstros := []Monstro{
		{Nome:"Esqueleto de lodo",Tipo:"Morte",Vida:60},
		{Nome:"Zumbi de sangue",Tipo:"Sangue",Vida:90},
		{Nome:"Existido",Tipo:"Conhecimento",Vida:30},
	}
	
	
	p1.Apresentar()
	fmt.Println("Você tem 2 opções escolha seu caminho com sabedoria")
	EscolherDirecao(&p1)
	if (p1.Andar == 1){
		fmt.Println("Olhe, está vendo? Um vilarejo ao fundo? Você pode conversar com algum aldeão em busca de missões")
	}else if p1.Andar == 2{
		inimigo := Aparecer(monstros)
		for inimigo.Vida > 0{
		fmt.Println("Realizando ataque")
		dano := verificarCritico(&p1, &inimigo)
		p1.Atacar(&inimigo, dano)
		fmt.Println("Vida atual do ",inimigo.Nome, "=",inimigo.Vida)}
		fmt.Println(inimigo.Nome,"Derrotado, Parabens Guerreiro!!!")
	}else {
		fmt.Println("Digite um valor válido")
	}
}
