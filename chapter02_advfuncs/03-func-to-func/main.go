package main

import (
	"fmt"
	"strings"
	"unicode"
)

type filterFunc func(int) bool

func main() {
	//signatures()
	fmt.Println("\n••• WITHOUT FUNC VALUES •••")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("evens      : %d\n", filterEvens(nums...))
	fmt.Printf("odds       : %d\n", filterOdds(nums...))

	fmt.Println("\n••• FUNC VALUES •••")
	fmt.Printf("evens      : %d\n", filter(isEven, nums...))
	fmt.Printf("odds       : %d\n", filter(isOdd, nums...))

	fmt.Println("\n••• MAPPING •••")
	fmt.Println(strings.Map(unpunct, "hello!!! HOW ARE YOU???? :))"))

}

func unpunct(r rune) rune {
	if unicode.IsPunct(r) {
		return -1
	}
	return unicode.ToLower(r)
}

func filter(f filterFunc, nums ...int) (filtered []int) {
	for _, v := range nums {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return
}

func filterOdds(nums ...int) (filtered []int) {
	for _, v := range nums {
		if isOdd(v) {
			filtered = append(filtered, v)
		}
	}
	return
}

func filterEvens(nums ...int) (filtered []int) {

	for _, v := range nums {
		if isEven(v) {
			filtered = append(filtered, v)
		}
	}
	return
}

func isEven(n int) bool {
	return n%2 == 0
}

func isOdd(m int) bool {
	return m%2 == 1
}
func signatures() {
	fmt.Println("••• FUNC SIGNATURES (TYPES) •••")
	fmt.Printf("isEven     : %T\n", isEven)
	fmt.Printf("isOdd      : %T\n", isOdd)
}
