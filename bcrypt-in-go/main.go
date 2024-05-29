package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	//Comparar as duas senhas, o cos é igual ao salt, é valor operacional, é o segundo parametro do generateFromPassword
	//minCost = 4
	//defaultCost = 10
	//maxCost = 34

	senha := "bruno123"
	senhaT := "bruno123"
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sb))

	senhaIgual := bcrypt.CompareHashAndPassword(sb, []byte(senhaT))
	if senhaIgual != nil {
		fmt.Println(senhaIgual)
	} else {
		fmt.Println("ok é a mesma senha")
	}

}