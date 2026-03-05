package main

import "fmt"

func main(){
	for {
		fmt.Println("Escolha a função a ser executada ");
		fmt.Println("Para Calcular Imc digite 1  ");
		fmt.Println("Para Calcular comissão digite 2 ");
		fmt.Println("Para calcular média escolar digite 3 ");
		fmt.Println("Para sair digite 4 ");
		var opcao int
		fmt.Scanln(&opcao)
		if opcao == 1 {

			var altura float64
			fmt.Println("Digite sua altura em metros ");
			fmt.Scanln(&altura)
			var peso float64
			fmt.Print("Digite seu peso kilogramas ");
			fmt.Scanln(&peso)
			var imc =   peso / (altura * altura)  

			fmt.Println("Seu indice de massa corpórea é = " ,imc);
		} else if opcao == 2 {
			var salario float64
			fmt.Println("Digite o valor do seu salario ");
			fmt.Scanln(&salario)

			var vendas float64
			fmt.Println("Digite o valor vendido no mês ");
			fmt.Scanln(&vendas)

			var comissao float64
			fmt.Println("Digite a porcentagem da sua comissao(Apenas numeros) ");
			fmt.Scanln(&comissao)
			var comissaoCal = (comissao / 100) * vendas
			var totalSalario = comissaoCal + salario

			fmt.Println("O valor da sua comissão é de  ",comissaoCal );
			fmt.Println("Salario total ",totalSalario);
		} else if opcao == 3 {
			var notaM float64
			fmt.Println("Digite o valor da nota mensal");
			fmt.Scanln(&notaM)

			var notaBm float64
			fmt.Println("Digite o valor da nota bimestral");
			fmt.Scanln(&notaBm)

			var media = (notaM + notaBm) / 2
			fmt.Println("Sua media é de",media );
			if media >= 6{
				fmt.Println("Aprovado ");
			} else {
				fmt.Println("Reprovado ");
			}
		}else if opcao == 4 {
			break
		}}}