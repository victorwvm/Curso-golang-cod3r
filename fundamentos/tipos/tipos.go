package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8(byte) uint16(short) uint32(int) uint64(long)

	var b byte = 255
	fmt.Println("Ö  byte e", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá meu nome é Victor Dantas"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	s2 := `Olá
	meu
	nome
	é
	Victor Dantas`

	fmt.Println(s2)
	fmt.Println("O tamanho da string é", len(s2))

	// char ???
	// var x = char 'b' ---- nao existe na linguagem
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
}
