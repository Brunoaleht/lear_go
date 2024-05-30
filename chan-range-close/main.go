package main

func main() {
	cannel := make(chan int)

	go myLoopChan(10, cannel) // goroutine para escrever no channel
	//não é necessário uma goroutine para ler o channel, eu tbm poderia ter feito uma goroutine para ler o channel, usando waitgroup para sincronizar
	// mas como o channel é fechado, não é necessário
	myReadChan(cannel)

}

func myLoopChan(total int, cs chan<- int) {
	for i := 0; i < total; i++ {
		cs <- i + 1
	}
	// close the channel, é importante para avisar que não terá mais valores
	close(cs)
}

func myReadChan(cr <-chan int) {
	// range é uma forma de ler os valores de um channel até que ele seja fechado, importante notar a diferencia entre o range para channel e o range para slice
	for v := range cr {
		println(v)
	}
}