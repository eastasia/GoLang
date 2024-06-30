package main

import "fmt"

func first() {
	msg := "Hello from the 1st function!"
	fmt.Println(msg)
}

func sqFive() {
	fmt.Printf("%v \n", 5*5)
}

func main() {

	first()
	fmt.Print("5 x 5 = ")
	sqFive()
}
