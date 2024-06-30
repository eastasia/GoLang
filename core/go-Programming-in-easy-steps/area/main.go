package main

import (
	"fmt"
	"math"
)

func main() {

	var rad, area, perim float64

	rad = 4
	fmt.Printf("\nRadius of Circle: %.2f \n", rad)

	area = math.Pi * (rad * rad)
	fmt.Printf("\nArea of Circle: %.2f \n", area)

	perim = 2 * (math.Pi * rad)
	fmt.Printf("\nPerimeter of Circle: %.2f \n", perim)

}
