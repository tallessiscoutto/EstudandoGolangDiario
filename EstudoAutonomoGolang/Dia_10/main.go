package main

import "fmt"


type Personagem struct {
	Nome string
	Vida int
}

func triplicar(x *int) {
	*x = *x * 3
}
func main(){
	num := 10
	triplicar(&num)
	fmt.Println(num)
}



