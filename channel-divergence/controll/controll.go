package controll

//controlando o número de goroutines para realizar o trabalho

import (
	"math/rand"
	"sync"
	"time"
)

func Controll() {
	canal1 := make(chan int)
	canal2 := make(chan int)
	number := 15
	goFuncs := 5

	go mandarCanal1(number, canal1)
	go outraParaCanal2(goFuncs, canal1, canal2)

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

func outraParaCanal2(n int, c1 , c2 chan int) {
	var wg sync.WaitGroup


	for i := 0; i < n; i++ { //esse for vai controlar o número de goroutines
		wg.Add(1)
		go func() {
			for v := range c1 {
				c2 <- trabalho(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(c2)
}

func trabalho(x int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return x
}