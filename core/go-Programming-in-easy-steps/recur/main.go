package main

import "fmt"

func main() {

	countDn(10)
}

func countDn(num int) {
	if num < 1 {
		fmt.Println("\t\t\t\tLift Off!")
	} else {
		fmt.Println("\t\tCountdown", num)
		num--
		countDn(num)
	}
}
