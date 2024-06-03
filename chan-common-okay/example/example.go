package example

func Example() {

	channel := make(chan int)

	go func() {
		channel <- 12
		close(channel)
	}()

	v, ok := <-channel
	println(v, ok) // o Ok = true é para informar q o valor realmente está neste canal, quando é Ok = false, significa que o canal está fechado e não tem mais valores para serem lidos

	v, ok = <-channel
	println(v, ok)

}