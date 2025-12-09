package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José Victor":    11325.43,
		"Gabriela Silva": 15542.42,
		"Pedro Junior":   1200.00,
	}

	funcsESalarios["Rafael Filho"] = 1350.0

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
