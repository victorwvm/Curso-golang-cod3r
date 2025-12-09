package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// Maps devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[18231237182] = "Pedro"
	aprovados[18123908213] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Buscar e deletar valor por uma chave
	fmt.Println(aprovados[12345678978])
	delete(aprovados, 12345678978)
	fmt.Println(aprovados[12345678978])
}
