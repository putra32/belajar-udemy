package main

import (
	"fmt"
	"sync"
	"testing"
)

func addToMap(value int, data *sync.Map, group *sync.WaitGroup) {
	defer group.Done()

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	var data = &sync.Map{}
	var group = &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go addToMap(i, data, group)
	}
	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
