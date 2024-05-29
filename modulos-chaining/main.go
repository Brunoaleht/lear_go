package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

func main() {
	users := []user{
		{"Bruno", 25, "Programador"},
		{"Lucas", 18, "Estagiario"},
		{"Gomes", 50, "Mestre"},
	}

	fmt.Println(users)
	//Enconder
	userSchema := json.NewEncoder(os.Stdout)
	err := userSchema.Encode(users)

	//Metodo chaning
	//err := json.NewEncoder(os.Stdout).Encode(users)

	if err != nil {
		fmt.Println(err)
	}
}