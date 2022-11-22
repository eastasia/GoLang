package main

import "fmt"

//place cursor at the struct field name and press Command + Shift + P ( Add Tags To Struct Fields)

type Table struct {
Name string `json:"name,omitempty"`
}


func main(){
	fmt.Println("Testing")
fmt.Println("Testing1")

}

func TestMe(p string){
	s := 0
	fmt.Println("This is a test function")
	fmt.Println(p,s)
	if len(p) > 5 {
		fmt.Println("Not Allowed")
	}
}