package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	var store list
	store = append(store, mobydick, &minecraft, &tetris)

	// #2
	store.discount(0.5)
	store.print()

	// #1
	//var p printer
	//p = &tetris
	//tetris.discount(.5)
	//p.print()
}
