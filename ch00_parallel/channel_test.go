package ch00_parallel

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestChannel01(t *testing.T) {
	inputChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		reqChan1 := make(chan int)
		go func() {
			defer func() {
				close(reqChan1)
				wg.Done()
			}()
			in := <-inputChan
			time.Sleep(time.Second * time.Duration(in))
			reqChan1 <- in
		}()

		reqChan2 := make(chan int)
		go func() {
			defer func() {
				close(reqChan2)
				wg.Done()
			}()
			in := <-inputChan
			time.Sleep(time.Second * time.Duration(in))
			reqChan2 <- in
		}()

		req1 := <-reqChan1
		req2 := <-reqChan2

		fmt.Printf("req1 = %v, req2 = %v\n", req1, req2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			inputChan <- rand.Intn(3)
		}
	}()

	wg.Wait()
}
