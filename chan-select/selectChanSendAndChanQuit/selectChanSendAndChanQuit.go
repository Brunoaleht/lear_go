package selectChanSendAndChanQuit

import "sync"

var wg sync.WaitGroup

func SelectChanSendAndChanQuit() {
	channel := make(chan int)
	channelQuit := make(chan int)

	wg.Add(2)

	go chanReceiveAndChanQuit(channel, channelQuit)
	go chanSendAndChanQuitSelect(channel, channelQuit)

	wg.Wait()
}

func chanReceiveAndChanQuit(cr chan int, cq chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		//Esse foi o erro q eu cometi na primeira vez fazendo: cr <- i, e não pode, pois estou tentando enviar um valor para um channel q não está esperando receber um valor
		value := <-cr
		println("Received:", value)
	}
	cq <- 0
}

func chanSendAndChanQuitSelect(cs chan int, cq chan int) {
	defer wg.Done()
	uniX := 1
	for {
		select {
		case cs <- uniX:
			uniX++
		case <-cq:
			return
		}
	}

	//Estou aproveitando o mecanismo do select q pode tanto receber um valor de um channel quanto enviar um valor para um channel
}