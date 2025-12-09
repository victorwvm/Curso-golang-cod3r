package main

import "fmt"

func main() {
	s := make([]int, 10) // Como array não foi criado, internamente o go criou um array interno e esse slice faz referência pra esse array

	s[9] = 12
	fmt.Println(s)

	// Slice com 10 elementos e 20 posições (crie um slice com 10 elementos, mas o array interno coloque 20)
	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s)) // Len = tamanho Cap = capacidade

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // Mesmo se atingir a capacidade maxima do array interno, quando adiciona novos elementos no slice ele internamente vai referênciando arrays maiores
	fmt.Println(s, len(s), cap(s))
}
