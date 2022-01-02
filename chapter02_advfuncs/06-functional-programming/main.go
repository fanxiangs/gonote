package main

import "fmt"

type filterFunc func(int) bool

func main() {
	dups := reverse(uniques())
	chained := chain(reverse(greater(1)), dups, uniques())
	fmt.Printf("dups chain : %v\n", filter(chained, 1, 1, 2, 3, 3, 3, 3, 0, 0))
}

func chain(filters ...filterFunc) filterFunc {
	return func(n int) bool {
		for _, next := range filters {
			if !next(n) {
				return false
			}
		}
		return true
	}
}
func uniques() filterFunc {
	seen := make(map[int]bool)

	return func(n int) (ok bool) {
		if !seen[n] {
			seen[n], ok = true, true
		}
		return
	}
}

func reverse(f filterFunc) filterFunc {
	return func(n int) bool {
		return !f(n)
	}
}

func greater(min int) filterFunc {
	return func(n int) bool {
		return n > min
	}
}
func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, n := range nums {
		if f(n) {
			filtered = append(filtered, n)
		}
	}
	return
}
