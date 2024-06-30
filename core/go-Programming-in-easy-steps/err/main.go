package main

import "fmt"

func main() {

	for i := 2; i >= -2; i-- {

		res, err := isPosInt(i)

		if err != nil {
			fmt.Println("Failed:", err)
		} else {
			fmt.Println(res, "Passed!")
		}
	}
}

func isPosInt(num int) (int, error) {

	if num < 1 {
		err := fmt.Errorf("%v not a positive integer", num)
		return -1, err
	}
	return num, nil
}
