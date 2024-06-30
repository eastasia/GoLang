package main

import "fmt"

func main() {

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				goto end
			}
			fmt.Println("Running i=", i, "j=", j)
		}
	}
end:
}
