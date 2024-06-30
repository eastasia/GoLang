package main

import "fmt"

func main() {

	_, b, c := cube(5)
	fmt.Println(b, "Cubed =", c)
}

func cube(num int) (string, int, int) {
	return "Result", num, (num * num * num)
}
