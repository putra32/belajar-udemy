package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type name struct {
	fistName   string
	middlename string
	lastName   string
}

func TestPool(t *testing.T) {
	pool := &sync.Pool{}

	boby := name{
		fistName:   "Boby",
		middlename: "Putra",
		lastName:   "Lubis",
	}

	pool.Put(boby)
	pool.Put(boby)
	pool.Put(boby)

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
