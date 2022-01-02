package main

import "fmt"

func main() {
	var (
		mobyDick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	var store list
	store = append(store, mobyDick, &minecraft, &tetris)
	store.print()

	fmt.Println(store[0] == mobyDick)
	fmt.Println(store[2] == &tetris)

}
