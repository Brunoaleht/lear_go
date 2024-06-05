package todd

import "sync"

func Todd() {
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)



	go send(par, impar)
	go receive(par, impar, converge)

	for v := range converge {
		println("Valor retirado em C", v)
	}

}

func send(p, i chan int) {
	x := 100
	for j := 0; j < x; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			i <- j
		}
	}
	close(p)
	close(i)
}

func receive(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}