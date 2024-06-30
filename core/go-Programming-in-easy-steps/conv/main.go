package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	str := "I have Nothing to declare except My Genius"
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.Title(strings.ToLower(str)))

	str = "  42  "
	fmt.Printf("\n%v Type: %T, Length: %v \n", str, str, len(str))

	str = strings.Trim(str, " ")
	fmt.Printf("\n%v Type: %T, Length: %v \n", str, str, len(str))

	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%v Type: %T \n", num, num)
	}
}
