package main

import "fmt"

func main() {
	//Criamos um canal com o tipo int
	c := make(chan int)

	//Aqui chamamos a função send e passamos o canal c, temos q por ao menos uma função em goRoutine para que o programa não fique esperando o canal ser lido
	go send(c)

	//Aqui chamamos a função receive e passamos o canal c
	receive(c)

	fmt.Println("About to exit")

}

// Quando queremos enviar um valor para um canal, usamos a  canal <-. Aqui vai ser a função send
func send(c chan<- int) {
	c <- 42
}

// Quando queremos receber um valor para de um canal, usamos a  <- canal. Aqui vai ser a função receive
func receive(c <-chan int) {
	value := <-c
	fmt.Println("O valor recebido do canal: ", value)
}