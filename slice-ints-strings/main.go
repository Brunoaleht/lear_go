package main

import (
	"fmt"
	"sort"

	"example.com/mod/slice-ints-strings/exercicio"
)

func main() {

	ss := []string{"x", "a", "b", "t", "z"}
	si := []int{5, 4, 100, 3, 98, 7, 1}

	fmt.Println("sem ordem", ss)

	sort.Strings(ss)
	sort.Ints(si)
	fmt.Println("String com sort", ss)
	fmt.Println("Int com sort", si)

	exercicio.Exercicio()
}