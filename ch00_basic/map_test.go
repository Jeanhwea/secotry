package ch00_basic

import "testing"

func TestMap01(t *testing.T) {
	passed := map[string]bool{}
	passed["key"] = true
	t.Logf("ok = %v", passed["key"])
	t.Logf("ok = %v", passed["key1"])
}
