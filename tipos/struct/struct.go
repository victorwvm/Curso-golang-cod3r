package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver
func (p produto) precoComDessconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lápis",
		preco:    1.50,
		desconto: 0.05,
	}
	fmt.Println(produto1, produto1.precoComDessconto())

	produto2 := produto{"Notebook", 2135.99, 0.15}
	fmt.Println(produto2.precoComDessconto())
}
