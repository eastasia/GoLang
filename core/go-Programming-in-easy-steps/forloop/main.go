package main

import "fmt"

func main() {

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop Iteration", counter)
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("\nOuter Loop Iteration", i)

		for j := 1; j <= 3; j++ {
			fmt.Println("\tInner Loop Iteration", j)
		}
	}
}
