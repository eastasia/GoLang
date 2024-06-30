package main

import "fmt"

func main() {

	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	more := make([]int, 3)
	more[0] = 200
	more[1] = 300
	more[2] = 400

	describe(nums)

	nums = append(nums, more[0])
	describe(nums)

	copy(nums[0:], nums[1:])
	nums = nums[:len(nums)-1]
	describe(nums)

}

func describe(nums []int) {
	fmt.Printf("\n%v Length: %v ", nums, len(nums))
	fmt.Printf("Capacity: %v \n", cap(nums))
}
