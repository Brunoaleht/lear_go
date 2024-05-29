package main

import "fmt"

func main() {
	channel := make(chan int)
	channel <- 34
	fmt.Println(<-channel) // deadlock isso ocorre pq não tenho um buffer para armazenar o valor, ou pq não tenho um receptor para receber o valor


	// channel := make(chan int , 1) // buffer de 1, se eu tentar enviar mais de 1 valor, ocorre deadlock
	// channel <- 34
	// channel <- 29
	// fmt.Println(<-channel) 
}