package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var contador = 0

func main() {
	goroutines := 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func1()
	}

	wg.Wait()
	fmt.Println(contador)

}

func func1() {
	for i := 0; i < 10; i++ {
		z := contador
		runtime.Gosched() //manda o computador executar outra goroutine e depois voltar
		z++
		contador = z
	}
	wg.Done()
}