package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Exercícios")
	// wg.Add(2)

	// go event("inicializando...")
	// go event("finalizando...")

	generatedGoroutines(50)

	wg.Wait()
}

// func event(event string ) {
// 	fmt.Println(event)
// 	wg.Done()
// }

func generatedGoroutines(i int){
	wg.Add(i)
	for j := 0; j < i; j++ {
	go func(i int) {
		fmt.Println("Goroutine", i+1)
		wg.Done()
	}(j)
	}
}


// criar duas goroutines e fazer com que elas sejam executadas antes do programa finalizar.
