package ch00_basic

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandom01(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 10000; i++ {
		go func() {
			for {
				num := rand.Intn(10000)
				t.Log(num)
			}
		}()
	}
	time.Sleep(30 * time.Second)
}
