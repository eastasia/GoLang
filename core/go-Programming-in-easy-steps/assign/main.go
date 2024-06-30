package main

import "fmt"

func main() {

	var a, b int

	a, b = 8, 4
	fmt.Println("Assigned Values:\ta =", a, "\tb =", b)

	a += b // a = 8 + 4.
	fmt.Println("Add and Assign:\t\ta =", a)
	a -= b // a = 12 - 4.
	fmt.Println("Subtract and Assign:\ta =", a)
	a *= b // a = 8 * 4.
	fmt.Println("Multiply and Assign:\ta =", a)
	a /= b // a = 32 / 4.
	fmt.Println("Divide and Assign:\ta =", a)
	a %= b // a = 8 % 4.
	fmt.Println("Remainder Assigned:\ta =", a)

}
