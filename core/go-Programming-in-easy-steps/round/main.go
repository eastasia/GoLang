package main

import (
	"fmt"
	"math"
)

func main() {

	var pi float64 = math.Pi
	fmt.Println("Pi:", pi)

	fmt.Println("\nFloor:", math.Floor(pi))
	fmt.Println("Ceiling:", math.Ceil(pi))
	fmt.Println("Round:", math.Round(pi))
	fmt.Println("Truncated:", math.Trunc(pi))

	fmt.Println("\nShort Pi:", math.Round(pi*100)/100)

	var e1 float64 = math.E
	fmt.Printf("\nE: %v %T \n", e1, e1)
	var e2 int = int(e1)
	fmt.Printf("Cast: %v %T \n", e2, e2)

}
