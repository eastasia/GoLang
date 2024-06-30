package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	c <- "Go Programming"

	msg := <-c
	fmt.Println(msg)
}
