package main

import "fmt"

func forArray() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Fori: arr at index %d is %d\n", i, arr1[i])
	}
	for i, v := range arr1 {
		fmt.Printf("Forr: arr at index %d is %d\n", i, v)
	}
}
func f(a [5]int) {
	a[1] = 2
	fmt.Println(a) // [1 2 3 4 5]
}
func fp(a *[5]int) {
	a[1] = 1000
	fmt.Println(a) // &[1 1000 3 4 5]
}

func pointerArray() {
	//var arr1 [3]int
	arr1 := [5]int{1, 2, 3, 4, 5}
	f(arr1)
	fmt.Println(arr1) // [1 2 3 4 5]
	fp(&arr1)
	fmt.Println(arr1) // [1 1000 3 4 5]
}
func main() {
	//forArray()
	pointerArray()
}
