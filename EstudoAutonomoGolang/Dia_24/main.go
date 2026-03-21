package main

import (
	"fmt"
)

type Notificador interface {
	Enviar(Mensagem string)
}
type Email struct {
}
type SMS struct {
}

func (e *Email) Enviar(Mensagem string) {
	fmt.Println(Mensagem)
}
func (s *SMS) Enviar(Mensagem string) {
	fmt.Println(Mensagem)
}

func notificar(n Notificador, Mensagem string) {
	n.Enviar(Mensagem)
}

func main() {
	e1 := Email{}
	s1 := SMS{}
	notificar(&e1, "Oi enviado por email")
	notificar(&s1, "Oi enviado por SmS")
}
