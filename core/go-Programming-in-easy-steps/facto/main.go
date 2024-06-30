package main

import "fmt"

func facto(num int) int {

	if num == 0 {
		return 1
	}
	return num * facto(num-1)
}

func main() {
	for i := 1; i <= 7; i++ {
		fmt.Println("Factorial", i, "=", facto(i))
	}
}
