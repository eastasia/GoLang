package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	txt := []byte("\nA thousand suns will stream on thee,\nA thousand moons will quiver.\n")

	err := ioutil.WriteFile("C:/Textfiles/Farewell.txt", txt, 0644)
	check(err)

	file, err := os.OpenFile("C:/Textfiles/Farewell.txt", os.O_APPEND, 0644)
	check(err)
	defer file.Close()

	slice := []byte("by Alfred Lord Tennyson.\n")
	nb, err := file.Write(slice)
	check(err)

	fmt.Printf("\nAppended: %v bytes - %v", nb, string(slice[:nb]))

}
