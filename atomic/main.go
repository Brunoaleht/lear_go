package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

var contador int64

func main() {
	goroutines := 5000
	func1(goroutines)
	wg.Wait()
	fmt.Println(contador)

}

func func1(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			atomic.LoadInt64(&contador)
			runtime.Gosched() //manda o computador executar outra goroutine e depois voltar
			wg.Done()
		}()
	}
}