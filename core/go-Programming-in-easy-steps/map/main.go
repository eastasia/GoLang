package main

import "fmt"

func main() {

	colors := make(map[string]string)

	colors["Red"] = "#FF0000"
	colors["Green"] = "#00FF00"
	colors["Blue"] = "#0000FF"

	fmt.Printf("\nColors: %v \n", colors)

	colors["Yellow"] = "#FFFF00"

	fmt.Printf("\nYellow Hex Code: %v \n", colors["Yellow"])

	delete(colors, "Yellow")

	for k, v := range colors {
		fmt.Printf("\nHex Code for %v is %v \n", k, v)
	}
}
