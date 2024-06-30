package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("\nEnter Text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		fmt.Println(scanner.Text())
	}

	fmt.Print("\nEnter Text:")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()

	words := strings.Fields(scanner.Text())

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else {
		for i, v := range words {
			fmt.Printf("%v: %v \n", i, v)
		}
	}

}
