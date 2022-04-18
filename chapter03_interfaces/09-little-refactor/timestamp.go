package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) string() string {
	if ts.IsZero() {
		return "unknown"
	}

	const layout = "2016/01"
	return ts.Format(layout) // same as: ts.Time.Format(layout)
}

// toTimestamp returns a timestamp value depending on the type of `v`.
func toTimestamp(v interface{}) (ts timestamp) {
	var t int

	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	}

	ts.Time = time.Unix(int64(t), 0)
	return ts
}
