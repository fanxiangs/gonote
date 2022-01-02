package main

import "fmt"

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}
	for _, it := range l {
		it.print()
	}
}
