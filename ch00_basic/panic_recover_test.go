package ch00_basic

import (
	"fmt"
	"testing"
	"time"
)

type PanicData struct {
	Code    int
	Message string
}

func (p *PanicData) String() string {
	return fmt.Sprintf("Code:%v, Message:%v", p.Code, p.Message)
}

func TestPanicRecover01(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recover: %v", r)
		}
	}()

	panic(&PanicData{Code: 1, Message: "Unknown Error"})
}

func Recover() {
	if r := recover(); r != nil {
		fmt.Printf("Recover: %v", r)
	}
}

func re() interface{} {
	return recover()
}

func TestPanicRecover02(t *testing.T) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		// panic(2)
		func() {
			panic(3)
		}()
	}()

	// panic(1)
	time.Sleep(time.Second)
}

func try(runner func(), catcher func(interface{})) {
	defer func() {
		if r := recover(); r != nil {
			catcher(r)
		}
	}()
	runner()
}

func TestPanicRecover03(t *testing.T) {
	try(func() {
		panic(121)
	}, func(r interface{}) {
		t.Log(r)
	})
}
