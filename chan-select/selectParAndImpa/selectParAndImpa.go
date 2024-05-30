package selectParAndImpa

func SelectParAndImpa() {
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
	close(cp)
	close(ci)
	cq <- true
}

func receivedValuesFromChannel(cp, ci chan int, cq chan bool) {
	for {
		select {
		case v := <-cp:
			println("Par:", v)
		case v := <-ci:
			println("Impar:", v)
		case <-cq:
			return
		}
	}
}

//Problema ao finalizar o meu loop, ele está sobrando um valor como se zero estivesse sendo enviado para o channel impa e par
//Ainda vou aprender a resolver isso
