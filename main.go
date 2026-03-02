//Nesse codigo estou aprendendo a aplicar a função retornando valor sem realizar a exibição direta, dividindo a tarefa com a main

package main

import "fmt"

func media(a, b, c float64) float64{
	resultado := (a + b + c) / 3
	return resultado
}

func main() {
	numero1 := 10.0
	numero2 := 8.0
	numero3 := 7.0

	mediaCalculada := media(numero1, numero2, numero3)
	fmt.Println("Media: ", mediaCalculada)
}