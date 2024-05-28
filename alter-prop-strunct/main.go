package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	x := pessoa{"Bruno", "Alexandre", 27}
	fmt.Println(x)
	mudeMe(&x)
	fmt.Println(x)
}

func mudeMe(p *pessoa) {
	(*p).nome = "Alexandre"
	p.sobrenome = "Bruno"
}