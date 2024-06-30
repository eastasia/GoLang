package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "I can resist everything except temptation"

	fmt.Printf("\nFound 'an': %v \n", strings.Contains(str, "an"))
	fmt.Printf("Found 'an' at: %v \n", strings.Index(str, "an"))

	fmt.Printf("Count of 'e': %v \n", strings.Count(str, "e"))

	fmt.Printf("Prefix 'ion': %v \n", strings.HasPrefix(str, "ion"))
	fmt.Printf("Suffix 'ion': %v \n", strings.HasSuffix(str, "ion"))

	fmt.Println(strings.Replace(str, "temptation", "chocolate", 1))

}
