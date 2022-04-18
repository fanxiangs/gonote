package main

import "fmt"

type Player struct{}

func (Player) Draw(power int) {
	fmt.Printf("Player is drawing %dkW of electrical power.\n", power)
}
