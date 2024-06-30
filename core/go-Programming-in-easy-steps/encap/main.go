package main

import (
	"cube"
	"fmt"
)

func main() {

	var box cube.Dims

	box.SetSize(2, 4, 6)

	fmt.Println("Footprint:", box.GetArea())
	fmt.Println("Volume:", box.GetVolume())

	// Below causes error as width is an unexported field.
	//fmt.Println("Width:", box.width)

	// Below causes error as getArea is an unexportd method.
	// fmt.Println("Area:", box.getArea())
}
