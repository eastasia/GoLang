package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	go count("Message", c)

	for {
		msg, open := <-c
		if !open {
			break
		}
		fmt.Println(msg)
	}
}

func count(msg string, c chan string) {

	for i := 0; i < 3; i++ {

		c <- msg + " at " + time.Now().Format("04:05")
		time.Sleep(1 * time.Second)
	}
	close(c)
}
