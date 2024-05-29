package main

import (
	"fmt"
	"sort"

	"example.com/mod/sort-customize/exercicio"
)

type car struct {
	name      string
	potencial int
	consumo   int
}

// ByPotencial
//Se para mim o importante é a menor potencia, basta mudar o sinal:   p[i].potencial "<" p[j].potencial

type ByPotencial []car

func (p ByPotencial) Len() int           { return len(p) }
func (p ByPotencial) Less(i, j int) bool { return p[i].potencial > p[j].potencial }
func (p ByPotencial) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ByConsumo
type ByConsumo []car

func (p ByConsumo) Len() int           { return len(p) }
func (p ByConsumo) Less(i, j int) bool { return p[i].consumo < p[j].consumo }
func (p ByConsumo) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//ByLucroParaPosto
//Neste caso vai dar mais lucro para o posto oq tiver o maior consumo, e por isso euq uero ele em cima

type ByLucroParaPosto []car

func (p ByLucroParaPosto) Len() int           { return len(p) }
func (p ByLucroParaPosto) Less(i, j int) bool { return p[i].consumo > p[j].consumo }
func (p ByLucroParaPosto) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	cars := []car{
		{"civic", 220, 15},
		{"gol", 110, 35},
		{"fusca", 225, 500},
	}
	fmt.Println("Normal", cars)
	sort.Sort(ByPotencial(cars))
	fmt.Println("Maior Potencia", cars)
	sort.Sort(ByConsumo(cars))
	fmt.Println("Menor consumo", cars)
	sort.Sort(ByLucroParaPosto(cars))
	fmt.Println("Maior Consumo", cars)

	exercicio.Exercicio()
}