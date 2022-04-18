package main

import (
	"fmt"
	"strconv"
	"time"
)

type book struct {
	title     string
	price     money
	published interface{}
}

func (b book) print() {
	p := format(b.published)
	fmt.Printf("%-15s: %s - (%v)\n", b.title, b.price.string(), p)
}

func format(v interface{}) string {
	if v == nil {
		return "unknown"
	}
	var t int

	if v, ok := v.(int); ok {
		t = v
	}

	if v, ok := v.(string); ok {
		t, _ = strconv.Atoi(v)
	}

	u := time.Unix(int64(t), 0)
	return u.String()
}
