package main

import "sync"

var wg sync.WaitGroup

func main() {
	
	channel := make(chan int)

	go generatedGoroutines(10, channel)

	for v := range channel {
		println("Valor retirado em C", v)
	}

}

func generatedGoroutines(i int, c chan int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func(x int) {
			for i := 0; i < x; i++ {
				c <- i
			}
			wg.Done()
		}(j)
		
	}
	wg.Wait()

	close(c)
}