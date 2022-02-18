package ch00_basic

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func TestJson02(t *testing.T) {
	obj := make(map[string]string)
	keyCount := 10000
	for i := 0; i < keyCount; i++ {
		key, val := fmt.Sprintf("k%08v", i), fmt.Sprintf("v%08v", i)
		obj[key] = val
	}

	bytes, _ := json.Marshal(obj)
	// os.WriteFile(fmt.Sprintf("/tmp/json-key-%08d.json", keyCount), bytes, 0644)

	str1 := strconv.Quote(string(bytes))
	os.WriteFile(fmt.Sprintf("/tmp/key-tidy-%05d.json", keyCount), []byte(str1), 0644)
}
