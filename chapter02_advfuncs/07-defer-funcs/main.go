package main

import (
	"fmt"
	"time"
)

func main() {
	//single()
	//stacked() // 先进后出
	findTheMeaning()
	//findTheMeaningNoDefer()

}

func findTheMeaningNoDefer() {
	start := time.Now()
	fmt.Printf("%s starts...\n", "findTheMeaning")
	time.Sleep(time.Second * 2)
	fmt.Printf("%s took %v\n", "findTheMeaning", time.Since(start))

}

func findTheMeaning() {
	defer measure("findTheMeaning")()
	time.Sleep(time.Second * 2)
}

func measure(name string) func() {
	start := time.Now()
	fmt.Printf("%s starts...\n", name)
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func stacked() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("the stacked func returns")
}

func single() {
	var count int
	defer func() {
		defer fmt.Printf("defer func count: %d\n", count)
	}()
	defer fmt.Printf("defer count: %d\n", count)
	count += 1
	fmt.Printf("single count: %d\n", count)
}
