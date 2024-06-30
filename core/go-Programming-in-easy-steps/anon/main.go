package main

import "fmt"

func main() {

	area := func(length, width int) int {
		return length * width
	}

	fmt.Printf("area Type: %T \n", area)
	fmt.Println("Area 1:", area(10, 4))
	fmt.Println("Area 2:", area(12, 5))

	counter := func() func() int {
		num := 0
		return func() int {
			num++
			return num
		}
	}()

	fmt.Printf("counter type: %T \n", counter)
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())
	fmt.Println("Count:", counter())

}
