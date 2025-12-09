package main

import "fmt"

func main() {
	x := 1
	y := 2

	// Apenas postfix (posfixada = depois do operando)
	x++ // x += 1 ou x = x + 1

	fmt.Println(x)

	y-- // y =- 1- ou y = y - 1
	fmt.Println(y)

	// fmt.Println(x == y--) Não pode misturar em expressões
}
