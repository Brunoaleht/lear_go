package selectLoopCase

func SelectLoopCase() {
	x := 6
	channelA := make(chan int)
	channelB := make(chan int)

	go func(j int) {
		for i := 0; i < j; i++ {
			channelA <- i + 1
		}

	}(x / 2)

	go func(j int) {
		for i := 0; i < j; i++ {
			channelB <- i + 1
		}

	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-channelA:
			println("Channel A:", v)
		case v := <-channelB:
			println("Channel B:", v)
		}
	}

	//isso é para demostrar q posso ter uma função q recebe vários channels e fazer um select com eles
}