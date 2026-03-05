package main

import (
	"fmt"
)
func main(){
	calcularVenda()
		}

func calcularVenda() {
	var nomeVendedor string
	fmt.Println("Digite nome do Vendedor ");
	fmt.Scanln(&nomeVendedor)
	var valorTotalVenda float64
	fmt.Print("Digite o valor Total da Venda  ");
	fmt.Scanln(&valorTotalVenda)
	var tipoCliente int
	fmt.Println("Digite tipo de Cliente sendo 1 = comum, 2 = vip , 3 = funcionario e 4 = Premium");
	fmt.Scanln(&tipoCliente)

	var comissao float64
	var tipoClientestring string
	switch tipoCliente {
	case 1:
		comissao = 0.05
		tipoClientestring = "comum"
	case 2:
		comissao = 0.10
		tipoClientestring = "VIP"
	case 3:
		comissao = 0
		tipoClientestring = "Funcionário"
	
	case 4:
		comissao = 0.20
		tipoClientestring = "Premium"
	}
		comissaoCalculada := valorTotalVenda * comissao
		fmt.Println("Vendedor :", nomeVendedor)
		fmt.Println("Valor da venda : ", valorTotalVenda)
		fmt.Println("Tipo de cliente :", tipoClientestring)
		fmt.Println("Comissão final :", comissaoCalculada)
}