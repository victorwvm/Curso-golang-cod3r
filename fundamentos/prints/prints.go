package main

import "fmt"

func main() {
	// Printa na mesama linha
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	// Começa na linha anterior (Porque o anterior nao quebrou a linha) no final quebra pra outra linha
	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	// fmt.Println("O valor de x é: " + x)

	// Go nao permite concatenar a string + o float, entao abaixo transformamos x em string pra cocatenar. Ou entao usamos a "," em vez do "+"
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é: ", x)

	// Tambem podemos usar o Printf(print formatado) %f = float %s = string
	fmt.Printf("O valor de x é: %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "Olá"

	// d = inteiro
	// f = float
	// 1.f = float com 1 casa decimal
	// t = boolean

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
