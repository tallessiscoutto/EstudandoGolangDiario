package main

import (
	"fmt"
)

type Notificador interface {
	Enviar()
}
type Email struct {
	mens string
}
type SMS struct {
	mens string
}

func (e *Email) Enviar() {
	fmt.Println(e.mens)
}
func (s *SMS) Enviar() {
	fmt.Println(s.mens)
}

func notificar(n Notificador) {
	n.Enviar()
}

func main() {
	e1 := Email{"Oi enviado por Email"}
	s1 := SMS{"Oi enviado por SMS"}
	notificar(&e1)
	notificar(&s1)
}
