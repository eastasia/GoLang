package main

import "fmt"

func main() {

	array := [3]string{"BMW", "Ford", "Opel"}

	slice := array[:]

	list(slice...)

	slice = append(slice, "Porsche", "Ferrari")

	list(slice...)

}

func list(autos ...string) {

	for i, v := range autos {
		fmt.Printf("\n%v. %v", i, v)
	}
	fmt.Println()
}
