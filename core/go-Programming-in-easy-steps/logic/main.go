package main

import "fmt"

func main() {

	var yes, no bool = true, false

	result := (yes && no)
	fmt.Println("AND Logic:\tyes && no\t", result)

	result = (yes || no)
	fmt.Println("OR Logic:\tyes || no\t", result)

	result = !yes
	fmt.Println("NOT Logic:\tyes =", yes, "\t!yes =", result)
}
