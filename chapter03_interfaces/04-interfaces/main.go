package main

import "fmt"

func main() {
	var (
		mobyDick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// 定义printer interface，可以打印列表中的值
	// 注意：每个添加的type需要有print method
	var store list
	store = append(store, mobyDick, &minecraft, &tetris)
	store.print()

	// interface的值是可以比较的
	fmt.Println(store[0] == mobyDick)
	fmt.Println(store[2] == &tetris)

}
