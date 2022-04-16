package main

import "fmt"

type Kettle struct{}

func (Kettle) Draw(power int) {
	fmt.Printf("Kettle is drawing %dkW of electrical power.\n", power)
}
