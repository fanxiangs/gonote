package main

import (
	"fmt"
)

type filterFunc func(int) bool

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("••• FUNC VALUES •••")
	fmt.Printf("evens      : %d\n", filter(isEven, nums...))
	fmt.Printf("odds       : %d\n", filter(isOdd, nums...))
	fmt.Printf("> 3        : %d\n", filter(greaterThan3, nums...))
	fmt.Printf("> 6        : %d\n", filter(greaterThan6, nums...))

	fmt.Println("\n••• CLOSURES •••")
	var min int
	greaterThan := func(n int) bool {
		return n > min
	}
	//min = 3
	//fmt.Printf("> 3        : %d\n", filter(greaterThan, nums...))
	var filterers []filterFunc
	for i := 1; i <= 4; i++ {
		current := 1
		filterers = append(filterers, func(n int) bool {
			min = current
			fmt.Printf("n!!!!!!!!!%d\n", n)
			return greaterThan(n)
		})
	}
	printer(filterers, nums...)
}

func printer(filterers []filterFunc, nums ...int) {
	for i, filterer := range filterers {
		fmt.Printf("> %d        : %d\n", i+1, filter(filterer, nums...))
	}
}

func greaterThan3(n int) bool { return n > 3 }
func greaterThan6(n int) bool { return n > 6 }

func isEven(n int) bool { return n%2 == 0 }
func isOdd(n int) bool  { return n%2 == 1 }

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, num := range nums {
		if f(num) {
			filtered = append(filtered, num)
		}
	}
	return
}
