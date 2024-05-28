package main

import "fmt"

func main() {

	returnWelcome(welcome)

}

func welcome() {
	fmt.Println("Olá, mundo")
}

func returnWelcome(f func()) {
	fmt.Println("olha eu aqui")
	f()
}
