package main

import "fmt"
type Personagem struct{
	Nome string;
	Vida int
	Tipo_ataque string
	Dano int}
type Monstro struct{
	Nome string
	Tipo string
	Vida int}
func (p *Personagem) Apresentar(){
	fmt.Println("Olá eu sou o ",p.Nome,"Seja bem vindo a essa jornada guerreiro, me ajude a derrotar os inimigos que aparecerem pelo caminho")
}
func Aparecer(m *Monstro, p *Personagem){
	fmt.Println(p.Nome,"Estava caminhando pela floresta quando do nada ouve um barulho em uns arbustos...")
	fmt.Println("Vou lá verificar, disse ele ")
	fmt.Println("Meu deus um ",m.Nome,"é hora do combate guerreiro")
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
	p1 := Personagem{Nome: "Talles",Vida:100,Tipo_ataque:"Sangue",Dano:15}
	m1 := Monstro{Nome:"Esqueleto de lodo",Tipo:"Morte",Vida:60}
	p1.Apresentar()
	Aparecer(&m1,&p1)
	for m1.Vida > 0{
		fmt.Println("Realizando ataque")
		Dano := verificarCritico(&p1, &m1)
		Atacar(Dano,&m1)
		fmt.Println(m1.Nome,"Recebeu",Dano,"de dano")
		fmt.Println("Vida atual do ",m1.Nome, "=",m1.Vida)}
		fmt.Println(m1.Nome,"Derrotado, Parabens Guerreiro!!!")
	}
