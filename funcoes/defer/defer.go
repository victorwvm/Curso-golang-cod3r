package main

import "fmt"

func obeterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obeterValorAprovado(6000))
	fmt.Println(obeterValorAprovado(3000))
}
