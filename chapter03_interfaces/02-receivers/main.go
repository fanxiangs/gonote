package main

func main() {
	var (
		mobyDick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	//传递minecraft的point给discount
	//类似于(&minecraft).discount(0.1)，编译器会自动隐式转换
	minecraft.discount(0.1)

	mobyDick.print()
	minecraft.print()
	tetris.print()
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
