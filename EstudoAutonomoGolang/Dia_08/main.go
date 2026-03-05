package main

import "fmt"


var Media float64
var Status string
var StatusAtualizado string
type Aluno struct{
	Nome string
	Nota1 float64
	Nota2 float64
}
func media_(a, b float64) float64{
	resultado := (a + b) / 2
	return resultado
}
func status(media float64)string {

	if media >= 7{
		Status = "Aprovado"
		
	} else {
		Status = "Reprovado"
	}
		return Status
}
func main(){
	alunos := []Aluno{
		{
			Nome: "Talles",
			Nota1: 7.6,
			Nota2: 8.4,
		},
		{
			Nome: "Renato",
			Nota1: 5.6,
			Nota2: 5.4,
		},
	}
	
	MediaCalculada := media_(alunos[0].Nota1, alunos[0].Nota2)
	StatusAtualizado = status(MediaCalculada)

	fmt.Println("Aluno: ", alunos[0].Nome)
	fmt.Println("Media: ", MediaCalculada)
	fmt.Println("Status: ", StatusAtualizado)

	fmt.Println("Alunos Cadastrados: ")
	for _, aluno := range alunos{
		fmt.Println(aluno.Nome)
	}
	
}



