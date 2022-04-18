package main

import "fmt"

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		it.print()
	}
}

// type assertion can extract the wrapped value.
// or: it can put the value into another interface.
func (l list) discount(ratio float64) {
	// you can declare an interface in a function/method as well.
	// interface is just a type.
	type discounter interface {
		discount(float64)
	}

	for _, it := range l {
		// you can assert to an interface.
		// and extract another interface.
		if it, ok := it.(discounter); ok {
			it.discount(ratio)
		}
	}
}
