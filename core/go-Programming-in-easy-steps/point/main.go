package main

import "fmt"

func main() {

	var num int = 20
	var ptr *int = &num

	fmt.Printf("num value: %v type: %T \n", num, num)
	fmt.Printf("num address: %v type: %T \n\n", ptr, ptr)

	fmt.Printf("num via pointer: %v type: %T \n", *ptr, *ptr)
	fmt.Printf("ptr address: %v type: %T\n\n", &ptr, &ptr)

	*ptr = 100
	fmt.Printf("new num value: %v type: %T \n", num, num)

}
