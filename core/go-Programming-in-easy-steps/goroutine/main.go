package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("\nLine-by-Line Execution...")
	count("moose")
	count("sheep")

	fmt.Println("\nConcurrent Execution...")
	go count("moose")
	count("sheep")
}

func count(item string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%v %v   ", i, item)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}
