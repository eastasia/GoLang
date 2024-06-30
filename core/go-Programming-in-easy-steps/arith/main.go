package main

import "fmt"

func main() {
	a := 8
	b := 4

	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Subtraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Remainder:\t", (a % b))

	a++
	fmt.Println("Increment:\t", a)
	b--
	fmt.Println("Decrement:\t", b)
}
