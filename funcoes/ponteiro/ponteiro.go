package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// Um ponteiro é representado por um *

func inc2(n *int) {
	// O * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// & usado para obter endereço da variável
	inc2(&n) // por referência
	fmt.Println(n)
}
