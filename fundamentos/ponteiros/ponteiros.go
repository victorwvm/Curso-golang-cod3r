package main

import "fmt"

func main() {
	i := 1

	/* Go não tem aritimética de ponteiros, mas pode compartilhar esse ponteiro com
	várias referências para que você possa ter um único espaço de memória senão referenciado várias várias
	*/
	var p *int = nil

	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando valor)
	i++

	fmt.Println(p, *p, i, &i)
}
