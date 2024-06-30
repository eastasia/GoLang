package main

import "fmt"

func main() {

	var cars [3]string
	cars[0] = "BMW"
	cars[1] = "Ford"
	cars[2] = "Opel"

	coords := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Cars:", cars)
	fmt.Println("Second Car:", cars[1])
	fmt.Println("X1,Y1:", coords[0][0])
	fmt.Println("X2,Y3:", coords[1][2])
}
