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

	txt, err := ioutil.ReadFile("C:/Textfiles/Oscar.txt")
	check(err)
	fmt.Println(string(txt))

	file, err := os.Open("C:/Textfiles/Oscar.txt")
	check(err)
	defer file.Close()

	pos, err := file.Seek(42, 0)
	check(err)

	slice := make([]byte, 15)
	nb, err := file.Read(slice)
	check(err)

	fmt.Printf("\n%v bytes @ %v: ", nb, pos)
	fmt.Printf("%v\n", string(slice[:nb]))

}
