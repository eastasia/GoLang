package main

import "fmt"

func main() {

	var zero, num, max int = 0, 0, 1
	var up, dn byte = 'A', 'a'

	result := (num == zero) // 0 == 0.
	fmt.Println("Equality:\tnum == zero\t", result)

	result = (up == dn) // A == a.
	fmt.Println("Equality:\tup == dn\t", result)

	result = (max != zero) //1 != 0.
	fmt.Println("Inequality:\tmax != zero\t", result)

	result = (zero > max) // 0 > 1.
	fmt.Println("Greater:\tzero > max\t", result)

	result = (max <= zero) // 1 <= 0.
	fmt.Println("Less or Equal:\tmax <= zero\t", result)

}
