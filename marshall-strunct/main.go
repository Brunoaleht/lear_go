package main

import (
	"encoding/json"
	"fmt"

	"example.com/mod/marshall-strunct/exercicio"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Saldo     float64
}

//Observação se eu quero exportar algo em Go tenho q por em Letra maiuscula no inicio

func main() {

	bruno := pessoa{
		Nome:      "Bruno",
		Sobrenome: "Alexandre",
		Idade:     28,
		Saldo:     25.25,
	}

	luke := pessoa{
		"luke",
		"Solsa",
		45,
		1000000.50,
	}

	b, err := json.Marshal(bruno)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	l, err := json.Marshal(luke)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(l))

	exercicio.Exercicios()
}