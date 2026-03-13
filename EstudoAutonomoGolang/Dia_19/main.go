package main

import "fmt"

type animal struct{
	especie string
	barulho string
}
type humano struct{
	nome string
	barulho string
}

type produzirSom interface {
	print()
}

func definirProdutorDeSom(p produzirSom){
	p.print()
}
func (h humano) print(){
	fmt.Println("Este ",h.nome,"fez um barulho através de um",h.barulho)
}
func (a animal) print(){
	fmt.Println("Este ",a.especie,"fez um barulho através de um",a.barulho)
}
func main() {
	animalNome := animal{"Cachorro","Latido"}
	humanos := humano{"Talles","Grito"}
	definirProdutorDeSom(animalNome)
	definirProdutorDeSom(humanos)

}
