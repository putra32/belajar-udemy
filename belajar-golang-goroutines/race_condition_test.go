package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestRaceConditon(t *testing.T) {
	var mutex = sync.Mutex{}
	var group = sync.WaitGroup{}

	x := 0
	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	group.Wait()

	fmt.Println("Counter: ", x)
}
