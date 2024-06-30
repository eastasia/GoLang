package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var num int = rand.Intn(20) + 1
	var guess int = 0
	var flag bool = true

	fmt.Print("\nGuess My Number 1-20: ")

	for flag {
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println(err)
		} else if guess > num {
			fmt.Print("Too High, Try Again: ")
		} else if guess < num {
			fmt.Print("Too Low, Try Again: ")
		} else if guess == num {
			fmt.Println("Correct - My Number Is", num)
			flag = false
		}
	}
}
