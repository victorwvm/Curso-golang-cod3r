package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"A": {
			"Arnaldo Dantas": 2100.00,
			"Ayrton Senna":   1614.13,
		},
		"B": {
			"Bruno Filho": 1413.00,
			"Brance Loyd": 2340.00,
		},
		"C": {
			"Cleidson Daniel": 4530.40,
			"Claudia Leite":   3420.00,
		},
	}

	delete(funcsPorLetra, "C")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
