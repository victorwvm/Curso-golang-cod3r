package main

import (
	"fmt"
	"time"
)

// Chanel (canal) - é a forma de comunicação entre goroutines
// Ele é sincrono, quando um dado não chega ele fica esperando o dado chegar

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)

	a, b, d := <-c, <-c, <-c // recebe dados do canal
	fmt.Println(a, b, d)

	// fmt.Println(<-c)
	// fmt.Println(<-c)

}
