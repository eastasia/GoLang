package main

import "fmt"

func main() {

	arr := [...]int{100, 200, 300, 400, 500}

	fmt.Println("No. of Elements:", len(arr))

	for i, v := range arr {
		arr[i] = v / 10
	}

	for i, v := range arr {
		fmt.Printf("Index: %v Value: %v \n", i, v)
	}
}
