package main

import (
	"fmt"

	"github.com/victorwvm/buscartitulo"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

// Mulktiplexar - misturar (mensagems) num canal

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		buscartitulo.Titulo("https://excalidraw.com/.br", "https://monkeytype.com/"),
		buscartitulo.Titulo("https://www.udemy.com.", "https://app.hub.la/"),
	)
	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)

}
