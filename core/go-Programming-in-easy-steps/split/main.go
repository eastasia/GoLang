package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "I can, and, I will."

	list("Fields", strings.Fields(str))

	list("Split", strings.Split(str, ","))
	list("SplitN", strings.SplitN(str, ",", 2))

	list("SplitAfter", strings.SplitAfter(str, ","))
	list("SplitAfterN", strings.SplitAfterN(str, ",", 2))

}

func list(desc string, subs []string) {
	fmt.Print("\n", desc, ": ")
	for _, v := range subs {
		fmt.Print("[", v, "] ")
	}
	fmt.Print("\n")
}
