package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// CUIDADO
	fmt.Println("Teste " + string(97)) // Se for concatenar inteiro precisa converter para string(no caso, aqui vai converter para o valor correspondente na tabela asc)

	// Int pra string
	fmt.Println("Teste " + strconv.Itoa(123))

	// String para int
	num, _ := strconv.Atoi("123") // Essa função retorna 2 valores: numero, erro (caso o valor nao
	// seja um int)
	fmt.Println("Teste", num)

	// Boolean
	b, _ := strconv.ParseBool("false")
	if b {
		fmt.Println("Verdadeiro")
	}

}
