package main

import "fmt"

type filterFunc func(int) bool

func main() {
	//signatures()
	funcValues()
}
func isEven(n int) bool {
	return n%2 == 0
}
func isOdd(n int) bool {
	return n%2 == 1
}

func signatures() {
	fmt.Println("••• FUNC SIGNATURES (TYPES) •••")
	fmt.Printf("isEven     : %T\n", isEven)
	fmt.Printf("isOdd      : %T\n", isOdd)
}

func funcValues() {
	fmt.Println("\n••• FUNC VALUES (VARS) •••")
	var value filterFunc
	fmt.Printf("value nil? : %t\n", nil == value)
	value = isEven
	fmt.Printf("value(2)   : %t\n", value(2))
	fmt.Printf("isEven(2)  : %t\n", isEven(2))
	fmt.Printf("isEven     : %p\n", isEven)
}
