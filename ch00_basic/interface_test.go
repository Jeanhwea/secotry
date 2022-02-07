package ch00_basic

import "testing"

type InterTest interface {
	GetName() string
}

type Human struct {
	name string
}

func (h *Human) GetName() string {
	return h.name
}

func TestInter01(t *testing.T) {
	var a InterTest
	if a != nil {
		a.GetName()
	}
}
