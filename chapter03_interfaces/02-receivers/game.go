package main

import "fmt"

type game struct {
	title string
	price float64
}

// 如果用一个类型的另一个方法的接受类型是point，
// 那么在其他方法也使用point
func (g *game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

// 定义接收类型为point的方法，这样就可以更新`price`的值
func (g *game) discount(ratio float64) {
	g.price *= 1 - ratio
}
