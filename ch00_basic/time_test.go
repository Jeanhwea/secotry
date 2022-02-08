package ch00_basic

import (
	"testing"
	"time"
)

func TestTime01(t *testing.T) {
	start := time.Now()
	time.Sleep(3 * time.Second)
	elapse := time.Now().Sub(start)
	t.Log(elapse)
}

func TestTime02(t *testing.T) {
	t.Logf("%v", time.Millisecond)
	t.Logf("%v", time.Microsecond)
	t.Logf("%v", time.Nanosecond)
}
