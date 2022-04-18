package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	// embed the product type into the book type.
	// all the fields and methods of the product will be available in this book type.
	product
	published interface{}
}

// book type's print method takes priority.
//
// + when you call it on a book value, the following method will
//   be executed.
//
// + if it wasn't here, the product type's print method would
//   have been executed.
func (b *book) print() {
	// the book can also call the embedded product's print method
	// if it wants to, as in here:
	b.product.print()

	p := format(b.published)
	fmt.Printf("\t - (%v)\n", p)
}

func format(v interface{}) string {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		return "unknown"
	}

	const layout = "2016/01"

	u := time.Unix(int64(t), 0)
	return u.Format(layout)
}
