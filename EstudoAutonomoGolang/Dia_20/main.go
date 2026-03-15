package main

import "fmt"

type cartaoCredito struct{
	
}
type pix struct{

}

type Pagamento interface {
	Pagar(valor float64)
}

func(c cartaoCredito) Pagar(valor float64){
	valor = 200
	fmt.Println("Pagamento efetuado com cartão de credito")
}

func(p pix)Pagar(valor float64){
	valor = 200
	fmt.Println("Pagamento efetuado com pix: Valor total: ",valor)
}

func processarPagamento(p Pagamento, valor float64){
	p.Pagar(valor)
}
func main() {
	c1 := cartaoCredito{}
	p1 := pix{}  
	
	processarPagamento(c1, 200)
	processarPagamento(p1, 200)
}
