package exercicio

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
	Job  string
}

func Exercicios() {
	users := []user{
		{"Bruno", 25, "Programador"},
		{"Lucas", 18, "Estagiario"},
		{"Gomes", 50, "Mestre"},
	}

	fmt.Println(users)

	sb, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))
}
