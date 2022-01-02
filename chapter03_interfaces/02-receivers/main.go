package main

func main() {
	var (
		mobyDick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	mobyDick.print()
	minecraft.discount(.1)
	minecraft.print()
	tetris.print()
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
