package main

import "fmt"

func main() {

	c := make(chan int)
	cs := make(chan<- int)
	cr := make(<-chan int)

	//Utilizando os canais e os tipos:
	fmt.Println("------------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cs\t%T\n", cs)
	fmt.Printf("cr\t%T\n", cr)

	//Aqui vamos converter o canal c para os tipos cs e cr
	fmt.Println("------------")
	fmt.Printf("c-cs\t%T\n", (chan<- int)(c))
	fmt.Printf("c-cr\t%T\n", (<-chan int)(c))

	fmt.Println("About to exit")

	//canal geral para um canal do tipo especifico funciona 
	//converter um canal especifico para geral não funcionar
	//converter um canal especifico para outro tipo de canal especifico não funcionar
	//não posso atribuir canais de tipos diferentes exemplo: c = cs, c = cr

}