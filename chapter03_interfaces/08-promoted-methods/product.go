package main

import "fmt"

type product struct {
	title string
	price money
}

type money float64

func (m money) string() string {
	return fmt.Sprintf("$%.2f", m)

}

func (p *product) print() {
	fmt.Printf("%-15s: %s\n", p.title, p.price.string())
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)

}
