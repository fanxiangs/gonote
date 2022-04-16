package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		blender Blender
		player  Player
		kettle  Kettle
	)
	socket := &Socket{50}
	fmt.Printf("Socket's available power is %d kW.\n", socket.power)
	if err := socket.Plug(blender); err != nil {
		fmt.Println("Blender cannot be powered up:", err)
	}
	if err := socket.Plug(player); err != nil {
		fmt.Println("Player cannot be powered up:", err)
	}
	if err := socket.Plug(kettle); err != nil {
		fmt.Println("Kettle cannot be powered up:", err)
	}
	fmt.Printf("Socket's available power is %d kW.\n", socket.power)

}
