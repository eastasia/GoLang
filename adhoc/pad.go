package main

import "fmt"

func main(){
	d := 123
	fmt.Printf("Padded: '%5d'\n", d)
	fmt.Printf("Padded: '%05d'\n", d)
}