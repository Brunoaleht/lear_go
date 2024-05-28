package main

import "fmt"

func main() {

	x := returnFunc()
	x()

}

func returnFunc() func() {
	return func() {
		fmt.Println("Hello, 世界")
	}
}