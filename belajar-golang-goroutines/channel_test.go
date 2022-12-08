package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannelGoroutines(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Boby Putra Lubis"
		fmt.Println("channel selesai mengirim data")
	}()

	// data := <-channel
	// fmt.Println(data)
	time.Sleep(5 * time.Second)
}

type address struct {
	city     string
	number   int32
	province string
	country  string
}

func PassByReference(addr *address) {
	addr.city = "Pariaman"
	addr.number = 135
}

func PastByRefence(n *int, r int) {
	*n = r
}

func TestPastByReferenceAndValue(t *testing.T) {
	// n := 15
	address1 := address{
		city:     "Padang",
		number:   120,
		province: "Sumatera Barat",
		country:  "Indonesia",
	}
	PassByReference(&address1)
	// PastByRefence(&n, 10)
	fmt.Println(address1)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Boby Putra Lubis"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyOut(channel)
	go OnlyIn(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Boby"
		channel <- "Putra"
		channel <- "Lubis"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "data ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima Data, ", data)
	}

	fmt.Println("Selesai")
}

func GiveMeResponse(channel chan string) {
	channel <- "Boby Putra"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1 ke ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel2 ke", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}
