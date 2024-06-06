package basic

import (
	"math/rand"
	"sync"
	"time"
)

func Basic() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	number := 15

	go mandarCanal1(number, canal1)
	go outraParaCanal2(canal1, canal2)

	for v := range canal2 {
		println(v)
	}
}

func mandarCanal1(x int, c chan int) {
	for i := 0; i < x; i++ {
		c <- i
	}
	close(c)
}

func outraParaCanal2(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			c2 <- trabalho(x)
		}(v)
	}
	//Aqui eu vou ter x goroutines, pois vai depender do número de elementos que eu tenho no canal c1
	wg.Wait()
	close(c2)
}

func trabalho(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return x
}