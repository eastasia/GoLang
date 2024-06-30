package main

import "fmt"

func main() {

	num := 100
	pi := 3.1415926536

	fmt.Printf("num: %v type:%T \n", num, num)
	fmt.Printf("num: %v type:%T \n\n", pi, pi)

	fmt.Printf("%%7d displays %7d \n", num)
	fmt.Printf("%%07d displays %07d \n\n", num)

	fmt.Printf("Pi is approximately %1.10f \n", pi)
	fmt.Printf("Right-aligned %20.3f rounded pi \n", pi)
	fmt.Printf("Left-aligned %-20.3f rounded pi  \n", pi)

}
