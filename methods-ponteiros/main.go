package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) FazerAniversario() {
	p.Age++
	fmt.Println(p.Age)
}

func (p *Person) Falar() {
	fmt.Println(p.Name, "diz oi")
}

type Human interface {
	Falar()
	FazerAniversario()
}

func Cumprimentar(h Human) {
	h.Falar()
}

func FazerAniversario(h Human) {
	h.FazerAniversario()
}

func main() {
	person1 := Person{"João", 20}

	Cumprimentar(&person1)
	FazerAniversario(&person1)

}