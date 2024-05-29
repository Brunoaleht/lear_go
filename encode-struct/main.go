package main

import (
	"encoding/json"
	"os"

	"example.com/mod/encode-struct/exercicio"
)

type pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	Saldo     float64
}

func main() {
	luke := pessoa{
		"luke",
		"Solsa",
		45,
		1000000.50,
	}
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(luke)

	exercicio.Exercicio()
}

//Diferencia de usar o marshal: o marchal eu tenho um intermediario q e salvar o valor em uma variavel primeiro para depois usar, o Encode não preciso salvar em uma variavel posso simplesment mandar o valor  para onde eu quiser serm varialvel