package main

import (
	"fmt"
)

func main(){
	var message interface{} = "Hello"
	s, ok := message.(string)
	if ok {
		fmt.Printf("%q %T\n", s, message)
	}


}