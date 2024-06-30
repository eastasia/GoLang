package main

import (
	"fmt"
)

func main() {

	nums := make(chan int)
	sqrs := make(chan int)

	go count(nums)
	go square(nums, sqrs)

	for i := 1; i <= 10; i++ {
		num := <-sqrs
		fmt.Printf("%v x %v = %v \n", i, i, num)
	}
}

func count(nums chan<- int) {
	for i := 1; i <= 10; i++ {
		nums <- i
	}
	close(nums)
}

func square(nums <-chan int, sqrs chan<- int) {
	for i := 1; i <= 10; i++ {
		num := <-nums
		sqrs <- num * num
	}
	close(sqrs)
}
