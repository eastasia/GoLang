package main

import (
	"flag"
	"fmt"
)

func main() {

	txt := flag.String("txt", "C#", "A string")
	num := flag.Int("num", 8, "An integer")
	sta := flag.Bool("sta", false, "A boolean")

	flag.Parse()

	fmt.Println("\nText:", *txt)
	fmt.Println("Number:", *num, " Status:", *sta)

}
