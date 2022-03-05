package ch00_parallel

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic01(t *testing.T) {
	var originCount int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				atomic.AddInt32(&originCount, 1)
			}
		}()
	}
	time.Sleep(5 * time.Second)
	curr := atomic.LoadInt32(&originCount)
	t.Logf("curr = %v", curr)
}

func TestAtomic02(t *testing.T) {
	var originCount int32 = 0
	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				originCount++
			}
		}()
	}
	time.Sleep(5 * time.Second)
	curr := originCount
	t.Logf("curr = %v", curr)
}
