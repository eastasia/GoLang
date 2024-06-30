package main

import "fmt"

func main() {

	var arr [5]int

	fmt.Println("No. of Elements:", len(arr))

	for i := 0; i < len(arr); i++ {
		arr[i] = i * i
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %v Value: %v \n", i, arr[i])
	}

}
