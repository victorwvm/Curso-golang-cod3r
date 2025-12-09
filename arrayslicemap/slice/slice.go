package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} // ARRAY
	s1 := []int{1, 2, 3}  // SLICE
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}

	// Slice não é um array! Slice define um pedaço de um array.

	s2 := a2[1:3]
	fmt.Println("Slice do elemento 1 ao 3 sem incluir o 3: ", a2, s2)

	// Slice não cria outro array, é apenas uma estrutura que tem um ponteiro pro primeiro elemento que ele aponta e um tamanho continuo...

	s3 := a2[:2]
	fmt.Println("Slice finalizando no elemento 2 sem inlcuir ele: ", a2, s3) // Novo slice, mas apontando pro mesmo array

	// Pode-se imaginar um slice como: tamanho e um ponteiro para um elemento de um array

	s4 := s2[:1]
	fmt.Println("Slice com slice: ", s2, s4)
}
