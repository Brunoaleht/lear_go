package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

var mu sync.Mutex

var contador = 0

func main() {
	goroutines := 10
	func1(goroutines)
	wg.Wait()
	fmt.Println(contador)

}

func func1(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			z := contador
			runtime.Gosched() //manda o computador executar outra goroutine e depois voltar
			z++
			contador = z
			mu.Unlock()
			wg.Done()
		}()
	}
}