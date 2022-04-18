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
	type discount interface {
		discount(float64)
	}
	for _, it := range l {
		if it, ok := it.(discount); ok {
			it.discount(ratio)
		}
	}
}

// // type assertion can pull out the real value behind an interface value.
// // it can also check whether the value convertable to a given type.
// func (l list) discount(ratio float64) {
// 	for _, it := range l {
// 		g, ok := it.(*game)
// 		if !ok {
// 			continue
// 		}

// 		g.discount(ratio)
// 	}
// }
