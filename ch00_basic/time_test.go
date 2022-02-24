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
	t.Logf("%v", time.Second)      // 1s = 1000ms
	t.Logf("%v", time.Millisecond) // 1ms = 1000us
	t.Logf("%v", time.Microsecond) // 1us = 1000ns
	t.Logf("%v", time.Nanosecond)  // 1ns
}

func TestTime03(t *testing.T) {
	t.Logf("%s", time.Now().Format(time.RFC3339))
	t.Logf("%s", time.Now().Format("2006-01-02 15:04:05.000"))

}

func TestTime04(t *testing.T) {
	a := time.Now().UnixNano()
	time.Sleep(1)
	b := time.Now().UnixNano()
	t.Logf("%v", int((b-a)/1000))
}
