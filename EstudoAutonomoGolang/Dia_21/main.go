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
	Defendendo bool
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
	return 
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
	fmt.Println("Olá caro viajante, meu nome é",aldeao,"Se estiver a procura de missões eu tenho uma boa pra você:")
	fmt.Println("Responda o aldeao, com o numero referente a instrução")
	fmt.Println("1 - Gostaria de uma missão")
	fmt.Println("2 - Olá tudo bem onde fica a taberna?")
	fmt.Println("3 - Não obrigado, to só passeando")
	var resposta int
	fmt.Scanln(resposta)
	aldeao = "Tadeu"
	switch resposta {
	case 1:
		fmt.Println("Tudo bem jovem, tem uma cabana na costa da montanha a uns 550 metros ao sul, lá você encontrará um rapaz que ta precisando de uma ajuda....")
		fmt.Println("Deseja ir até a montanha? 1 - Sim e 2 - Não")
		var resposta2 int
		fmt.Scanln(resposta2) 
		if (resposta2 == 1) {
			fmt.Println("Indo até a montanha")
			fmt.Println("Caraca essa cabana ta longe ta longe em ")
		}else {
			fmt.Println("Tudo bem, então não tenho missões pra você")
		}
		return
	}

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

func combate(p *Personagem, m *Monstro){
		for m.Vida > 0 && p.Vida > 0{
			p.Defendendo = false
			fmt.Println("Turno do jogador")
			acao = escolherAcao()
			switch acao {
			case 1:
				fmt.Println("Realizando ataque")
				dano := verificarCritico(p, m)
				Atacar(m, dano)
				fmt.Println("O ",m.Nome,"Sofreu",dano)
				fmt.Println("Vida atual do ",m.Nome,"é ",m.Vida)
			case 2:
				fmt.Println("Você está com ",p.Vida,"de vida")			
				fmt.Println("E entrou em posição de defesa")
				p.Defendendo = true

			case 3:
				fmt.Println("Você fugiu da batalha")
				return
			default:
				fmt.Println("Ação Invalida")
				continue
			}

			if m.Vida > 0 {
				fmt.Println("\nO",m.Nome,"Atacaou")
				Atacar(p, m.Dano)
				fmt.Println("Dano causado",m.Dano)
				fmt.Println("Sua vida atual é de ",p.Vida)
			}
	}
		if p.Vida <= 0 {
			fmt.Println("\nVocê foi derrotado...")

		}else{
			fmt.Println("\nMonstro Derrotado")
		}
}
func (m *Monstro)ReceberDano(dano int){
		m.Vida -= dano
	
}
func (p *Personagem)ReceberDano(dano int){
		if p.Defendendo{
		fmt.Println("Defesa ativada,Dano reduzido")
		dano = dano / 2
		p.Defendendo = false
		}
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
	EscolherDirecao(&p1)
	switch p1.Andar {
	case 1:
		vilareijo(&p1, aldeao)
	
	case 2:
		inimigo := Aparecer(monstros)
		fmt.Println("Um ",inimigo.Nome,"apareceu!")
		combate(&p1, &inimigo)
}
	

	
}
