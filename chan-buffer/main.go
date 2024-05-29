package main

import "fmt"

func main() {
	channel := make(chan int , 1)
	channel <- 34
	fmt.Println(<-channel) //com buffer não ocorre deadlock, buffer é um espaço de armazenamento temporário

}