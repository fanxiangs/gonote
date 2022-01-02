package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 7}
	fmt.Printf("nums: %d\n", nums)

	n := avgNoVariadic(nums)
	fmt.Printf("avgNoVariadic: %d\n", n)

	n = avg(2, 3, 7)
	fmt.Printf("avg(2, 3, 7): %d\n", n)

	// use ... to pass a slice
	n = avg(nums...)
	fmt.Printf("avg(nums...): %d\n", n)

	double(nums...)
	fmt.Printf("double(nums...) : %d\n", nums)

	fmt.Printf("\nmain.nums: %p\n", nums)
	investigate("passes main.nums", nums...)
	investigate("passes args", 4, 6, 14)
	investigate("no args")

}

func investigate(msg string, nums ...int) {
	fmt.Printf("investigate.nums: %12p  ->  %s\n", nums, msg)
	if len(nums) > 0 {
		fmt.Printf("\tfirst element: %d\n", nums[0])
	}

}

func double(nums ...int) {
	for _, v := range nums {
		v *= 2
	}
}

func avg(nums ...int) int {
	return sum(nums) / len(nums)
}

func avgNoVariadic(nums []int) int {
	return sum(nums) / len(nums)
}

func sum(nums []int) (total int) {
	for _, num := range nums {
		total += num
	}
	return
}
