package main

import (
	"fmt"
)

var Tipo_ataque int

type Acao interface {
	Atacar()
	Especial()
}
type Guerreiro struct {
	Nome string
	Dano int
}
type Mago struct {
	Nome string
	Mana int
}
type Arqueiro struct {
	Nome    string
	Flechas int
}

func (g *Guerreiro) Atacar() {
	fmt.Println("Você atacou com uma espada")
	g.Dano = 30
	fmt.Println("Dano causado", g.Dano)
}
func (g *Guerreiro) Especial() {
	fmt.Println("Você usou a habilidade lamina sangrenta")
	g.Dano = 50
	fmt.Println("Dano causado", g.Dano)
}
func (m *Mago) Atacar() {
	fmt.Println("Você atacou com um cajado")
	m.Mana = 30
	fmt.Println("Mana causado", m.Mana)
}
func (m *Mago) Especial() {
	fmt.Println("Você usou a habilidade Carnificina")
	m.Mana = 50
	fmt.Println("Mana gasta", m.Mana)
}
func (a *Arqueiro) Atacar() {
	fmt.Println("Você disparou com o arco")
	a.Flechas = 30
	fmt.Println("Dano causado", a.Flechas)
}
func (a *Arqueiro) Especial() {
	fmt.Println("Você usou a habilidade chuva de flechas")
	a.Flechas = 50
	fmt.Println("Dano causado", a.Flechas)
}
func personagem(a Acao) {
	a.Atacar()
	a.Especial()
}
func main() {
	g1 := Guerreiro{"Talles", 100}
	m1 := Mago{"Jonas", 100}
	a1 := Arqueiro{"Natan", 100}
	personagens := []Acao{&g1, &m1, &a1}

	for _, p := range personagens {
		personagem(p)
	}
}
