package main

import (
	"fmt"
)

func main() {

	c := make(chan string, 10)

	c <- "Go Programming"

	msg := <-c
	fmt.Printf("\n%v \n\n", msg)

	c <- "Go Programming"
	c <- "in "
	c <- "easy "
	c <- "steps "

	close(c)

	for msg := range c {
		fmt.Println(msg)
	}

}
