package main

import (
	"encoding/json"
	"fmt"
)

type informacao struct {
	Nome      string  `json:"Nome"`
	Sobrenome string  `json:"Sobrenome"`
	Idade     int     `json:"Idade"`
	Saldo     float64 `json:"Saldo"`
}

// `json:"Nome"`, isso são tags, observação olhar a documentação para saber usar melhor
// tags são usadas para dizer o nome do que está chegando e do que eu estou mandando exemplo:
//type informacao struct {
//Nome      string  `json:"name"`
//Sobrenome string  `json:"lastName"`
//Idade     int     `json:"age"`
//Saldo     float64 `json:"count"`
//}
//estou recebendo do meu json em minusculo e em ingles, e no meu codigo estou usando em portugues com a primeira letra maiuscula para poder export para json

func main() {

	sliceByte := []byte(`{"Nome":"Bruno","Sobrenome":"Alexandre","Idade":28,"Saldo":25.25}`)
	var bruno informacao

	err := json.Unmarshal(sliceByte, &bruno)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(bruno)
	fmt.Println(bruno.Idade)
}