package main

func main() {

	xNumber := 30
	channelPar := make(chan int)
	channelImpar := make(chan int)
	channelQuit := make(chan bool)

	go sendNumberToChannel(xNumber, channelPar, channelImpar, channelQuit)
	receivedValuesFromChannel(channelPar, channelImpar, channelQuit)

}

func sendNumberToChannel(number int, cp, ci chan int, cq chan bool) {
	for i := 0; i < number; i++ {
		if i%2 == 0 {
			cp <- i
		} else {
			ci <- i
		}
	}
	cq <- true
}

func receivedValuesFromChannel(cp, ci chan int, cq chan bool) {
	for {
		select {
		case v := <-cp:
			println("Par:", v)
		case v := <-ci:
			println("Impar:", v)
		case v, ok := <-cq:
			if !ok {
				println("Channel Quit Ferrou:", v)
			} else {
				println("Channel Quit Encerrando:", v)
			}
			return
		}
	}
}