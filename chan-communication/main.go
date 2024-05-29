package main

import "fmt"

func main() {
	channel := make(chan int)
	go func() {
		channel <- 37
	}()
	
	fmt.Println(<-channel) 

}