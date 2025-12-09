package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3 // Infrerência de tipo
	i += 3 // i = i + 3
	i -= 3 // i = i - 3
	i /= 2 // i = i / 2
	i *= 2 // i = i * 2
	i %= 2 // i = i % 2

	fmt.Println(i)

	x, y := 1, 2
	fmt.Println(x, y)

	// Trocou valores de 2 variaveis. X passou a ser Y e Y passou a ser X
	x, y = y, x
	fmt.Println(x, y)
}
