package main

import (
	"fmt"
	"math/rand"
)
const (
	atacar = 1
	defender = 2
	recuar = 3
)

var acao int
var aldeao string
var turno int
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
	Vida int
	Dano int
}
type Alvo interface{
	ReceberDano(dano int)
}
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

func Atacar(a Alvo, dano int){
	a.ReceberDano(dano)
}

func escolherAcao()int{
	fmt.Println("Escolha sua ação:")
	fmt.Println("1 = atacar")
	fmt.Println("2 = Defender")
	fmt.Println("3 = Recuar")
	var acao int
	fmt.Scanln(&acao)
	return acao
}

func combate(p *Personagem, monstros []Monstro){
		inimigo := Aparecer(monstros)
		for inimigo.Vida > 0 && p.Vida > 0{
			fmt.Println("Turno do jogador")
			acao = escolherAcao()
			switch acao {
				
			}
			for turno = 0; turno < 99; turno++ {
				escolherAcao(p, a, dano)
				fmt.Println("Realizando ataque")
				dano := verificarCritico(p, &inimigo)
				Atacar(&inimigo, dano)
				fmt.Println("Vida atual do ",inimigo.Nome, "=",inimigo.Vida)
				Atacar(p ,inimigo.Dano)
				fmt.Println("Vida atual de",p.Nome,"=", p.Vida)
	}
	}
		if p.Vida == 0{
			fmt.Println("Que pena, você foi derrotado")
			return
		} else if inimigo.Vida <= 0{
			fmt.Println(inimigo.Nome,"Derrotado, Parabens Guerreiro!!!")
		}

}
func (m *Monstro)ReceberDano(dano int){
		m.Vida -= dano
	
}
func (p *Personagem)ReceberDano(dano int){
		p.Vida -= dano
}
func main(){
	p1 := Personagem{Nome: "Talles",Vida:100,Tipo_ataque:"Sangue",Dano:15,}
	monstros := []Monstro{
		{Nome:"Esqueleto de lodo",Tipo:"Morte",Vida:60, Dano: 30},
		{Nome:"Zumbi de sangue",Tipo:"Sangue",Vida:90, Dano: 50},
		{Nome:"Existido",Tipo:"Conhecimento",Vida:30, Dano: 10},
	}
	p1.Apresentar()
	fmt.Println("Você tem 2 opções escolha seu caminho com sabedoria")
	EscolherDirecao(&p1)
	if(p1.Andar == 1){
		fmt.Println("Olhe, está vendo? Um vilarejo ao fundo? Você pode conversar com algum aldeão em busca de missões")
		vilareijo(&p1, aldeao)
	}else if p1.Andar == 2{
		combate(&p1, monstros)
	}else {
		fmt.Println("Digite um valor válido")
	}
}
