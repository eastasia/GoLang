package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str string = "42"
	var char byte = 'A'

	num, err := strconv.Atoi(str)
	fmt.Printf("num:%v %T %v \n", num, num, err)

	decimal := float64(num)
	fmt.Printf("decimal: %.2f %T \n", decimal, decimal)

	fmt.Printf("char:%v %T %v \n", char, char, string(char))

}
