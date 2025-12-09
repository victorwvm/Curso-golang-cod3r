package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// Jsojn pra struct
	var p2 produto
	jsonString := `{"id:2", "nome:Celular", "preco: 677.90", "Tags:{"Eletronicos", "Smartphone"}}`
	json.Unmarshal([]byte(jsonString), &p2)
	// fmt.Println(p2.Tags[1])
}
