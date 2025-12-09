package main

import (
	"fmt"
	"time"

	"github.com/victorwvm/buscartitulo"
)

func oMaisRapido(url1, url2, url3 string) string {
	c1 := buscartitulo.Titulo(url1)
	c2 := buscartitulo.Titulo(url2)
	c3 := buscartitulo.Titulo(url3)

	// estrutura de controle especifica pra concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda"
	}
}

func main() {
	campeao := oMaisRapido(
		"https://monkeytype.com/",
		"https://excalidraw.com/",
		"https://www.bratgenerator.com/",
	)
	fmt.Println(campeao)
}
