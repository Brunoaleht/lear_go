package main

import (
	"encoding/json"
	"fmt"

	"example.com/mod/json-unmarshall-struct/exercicio"
)

type informacao struct {
	Nome      string  `json:"Nome"`
	Sobrenome string  `json:"Sobrenome"`
	Idade     int     `json:"Idade"`
	Saldo     float64 `json:"Saldo"`
}

func main() {

	sliceByte := []byte(`{"Nome":"Bruno","Sobrenome":"Alexandre","Idade":28,"Saldo":25.25}`)
	var bruno informacao

	err := json.Unmarshal(sliceByte, &bruno)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(bruno)
	fmt.Println(bruno.Idade)

	exercicio.Exercicio()
}