package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	nums := rand.Perm(59)

	for i := 0; i < len(nums); i++ {
		nums[i]++
	}

	str := "\nYour Six Lucky Numbers: "
	for i := 0; i < 6; i++ {
		str += strconv.Itoa(nums[i])
		if i != 5 {
			str += " - "
		}
	}

	fmt.Println(str)

}
