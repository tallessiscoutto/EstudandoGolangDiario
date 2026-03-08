package main

import "fmt"

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
	p1 := Personagem{Nome: "Talles",Vida:100,Tipo_ataque:"Sangue",Dano:15,}
	fmt.Println("Olá eu sou o ",p1.Nome,"Seja bem vindo a essa jornada guerreiro, me ajude a derrotar os inimigos que aparecerem pelo caminho")
}
func EscolherDirecao(m *Monstro, p *Personagem){
	fmt.Println("Escolha uma direção 1 = Esquerda e 2 = Direita")
	fmt.Scanln(&p.Andar)
	if (p.Andar == 1){
		fmt.Println("Olhe, está vendo? Um vilarejo ao fundo? Você pode conversar com algum aldeão em busca de missões")
	}else if p.Andar == 2{
		Aparecer(m, p)
	}else {
		fmt.Println("Digite um valor válido")
	}
}
func Aparecer(m *Monstro, p *Personagem){
	m1 := Monstro{Nome:"Esqueleto de lodo",Tipo:"Morte",Vida:60}
	fmt.Println("Cuidado, esse caminho é meio sombrio")
	fmt.Println("Arbustos se mexem ")
	fmt.Println("Meu deus um ",m1.Nome,"é hora do combate guerreiro")
	return
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
	return p.Dano}
func Atacar(Dano int, m *Monstro)int{
	m.Vida -= Dano
	return m.Vida}
func main(){
	
	
	EscolherDirecao()
	Apresentar()
	Aparecer()
	for m1.Vida > 0{
		fmt.Println("Realizando ataque")
		Dano := verificarCritico(&p1, &m1)
		Atacar(Dano,&m1)
		fmt.Println(m1.Nome,"Recebeu",Dano,"de dano")
		fmt.Println("Vida atual do ",m1.Nome, "=",m1.Vida)}
		fmt.Println(m1.Nome,"Derrotado, Parabens Guerreiro!!!")
	}
