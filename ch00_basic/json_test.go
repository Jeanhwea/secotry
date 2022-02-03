package ch00_basic

import (
	"encoding/json"
	"testing"
)

type Book struct {
	Name      string  `json:"Name"`
	PageCount int     `json:"PageCount"`
	Price     float64 `json:"Price"`
}

func TestJson01(t *testing.T) {
	book01 := &Book{
		Name:      "A Litte Store",
		PageCount: 238,
		Price:     43.9,
	}
	str, _ := json.Marshal(book01)
	t.Logf("str = %s", str)
}
