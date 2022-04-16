package main

func main() {
	var (
		//mobyDick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	var items []*game
	items = append(items, &minecraft, &tetris)
	my := list(items)
	// 即使my为nil也可以调用方法
	//my = nil
	my.print()
}
