package main

import "fmt"

type huge struct {
	// about ~200mb
	games [10_000_000]game

	//games [10]game
}

func (h *huge) addr() {
	fmt.Printf("%p\n", h)
}
