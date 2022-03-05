package ch00_basic

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
)

type Book struct {
	Name       string   `json:"Name"`
	PageCount  int      `json:"PageCount"`
	Price      float64  `json:"Price"`
	AuthorList []string `json:"Author,omitempty"`
}

func TestJson01(t *testing.T) {
	book01 := &Book{
		Name:       "A Litte Store",
		PageCount:  238,
		Price:      43.9,
		AuthorList: []string{"Tom"},
	}
	str, _ := json.Marshal(book01)
	t.Logf("str = %s", str)
}

func prettifyJson(src []byte) (dest string) {
	var obj map[string]interface{}
	json.Unmarshal(src, &obj)
	newBytes, _ := json.MarshalIndent(obj, "", "  ")
	dest = string(newBytes)
	return
}

type Dog struct {
	Name        string `json:"name,omitempty"`
	Age         int64  `json:"age,omitempty"`
	Description string `json:"description,omitempty"`
	JsonIgnore  string `json:"-"`
}

func TestJson02(t *testing.T) {
	str01 := `{"name":"Scott"}`
	scott := &Dog{}

	json.Unmarshal([]byte(str01), scott)
	scott.Age = 8
	scott.JsonIgnore = "ignored field"
	t.Logf("scott = %+v", scott)
	t.Logf("scott = %+#v", scott)

	data, _ := json.Marshal(scott)
	t.Log(prettifyJson(data))
}

func TestJson10(t *testing.T) {
	obj := make(map[string]string)
	keyCount := 10000
	for i := 0; i < keyCount; i++ {
		key, val := fmt.Sprintf("k%08v", i), fmt.Sprintf("v%08v", i)
		obj[key] = val
	}

	bytes, _ := json.Marshal(obj)
	// os.WriteFile(fmt.Sprintf("/tmp/json-key-%08d.json", keyCount), bytes, 0644)

	str1 := strconv.Quote(string(bytes))
	os.WriteFile(fmt.Sprintf("/tmp/key-tidy-%05d.txt", keyCount), []byte(str1), 0644)
}

func saveQuote(name string, obj interface{}) {
	data, _ := json.Marshal(obj)
	quoted := strconv.Quote(string(data))
	os.WriteFile(name, []byte(quoted), 0644)
}

func TestJson11(t *testing.T) {
	extra := make(map[string]interface{})
	for i := 1; i <= 10; i++ {
		key := fmt.Sprintf("key%04d", i)
		val := fmt.Sprintf("val%04d", i)
		extra[key] = val
	}
	saveQuote("/tmp/data.txt", extra)
}
