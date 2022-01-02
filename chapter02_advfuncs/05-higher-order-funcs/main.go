package main

import "fmt"

type filterFunc func(int) bool

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd := reverse(reverse(isEven))
	fmt.Printf("reversed   : %t\n", odd(8))
	fmt.Printf("> 3        : %d\n", filter(greater(3), nums...))
}

func greater(min int) filterFunc {
	return func(n int) bool {
		return n > min
	}
}

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, num := range nums {
		if f(num) {
			filtered = append(filtered, num)
		}
	}
	return
}

func reverse(f filterFunc) filterFunc {
	return func(n int) bool {
		return !f(n)
	}
}

func isEven(n int) bool { return n%2 == 0 }
func isOdd(m int) bool  { return m%2 == 1 }
