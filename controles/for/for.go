package main

import (
	"fmt"
	"time"
)

func main() {

	// Contador (While)
	i := 1
	for i <= 10 { // Não tem while em GO
		fmt.Println(i)
		i++
	}

	// For tradicional (1 inicialização, 2 condição, 3 incremento)
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	// For aninhado
	fmt.Println("\nMisturando....")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	for {
		// Laço infinito
		fmt.Println("Para sempre....")
		time.Sleep(time.Second)
	}
}
