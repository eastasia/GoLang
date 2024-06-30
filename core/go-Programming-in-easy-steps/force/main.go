package main

import "fmt"

func main() {

	sum := 2*3 + 4 - 5
	fmt.Printf("Default Order: %v\n", sum)

	sum = 2 * ((3 + 4) - 5)
	fmt.Printf("Forced Order: %v \n\n", sum)

	sum = 7 % 3 * 2
	fmt.Printf("Default Order: %v \n", sum)

	sum = 7 % (3 * 2)
	fmt.Printf("Forced Order: %v \n\n", sum)

}
