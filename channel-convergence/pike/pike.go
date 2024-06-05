package pike

import (
	"fmt"
	"math/rand"
	"time"
)

func Pike() {

	umaVariavel := converge(trabalho("primeiro"), trabalho("segundo"))

	for i := 0; i < 10; i++{
		fmt.Println(<-umaVariavel)
	}

}

func trabalho(msg string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %s diz: %d", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(msg, canal)

	return canal
}

func converge(a, b chan string) chan string {
	newCanal := make(chan string)

	go func() {
		for {
			newCanal <- <-a
		}
	}()

	go func() {
		for {
			newCanal <- <-b
		}
	}()

	return newCanal
}