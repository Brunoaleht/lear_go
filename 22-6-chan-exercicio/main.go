package main

import "sync"

func main() {
	channel := make(chan int)
	converge := make(chan int)

	go send(channel)
	go receive(channel, converge)

	for v := range converge {
		println("Valor retirado em C", v)
	}

}

func send(c chan<- int) {
	number := 100

	for i := 0; i < number; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int, converge chan<- int) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range c {
			converge <- v

		}
		wg.Done()
	}()
	wg.Wait()
	close(converge)
}