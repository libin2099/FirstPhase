package main

import (
	"fmt"
	"time"
)

func produce(channel chan int) {
	for i := 1; i <= 10; i++ {
		channel <- i
		time.Sleep(50 * time.Millisecond)
	}
	close(channel)
}

func consume(channel chan int) {
	for {
		curNum, ok := <-channel
		if !ok {
			return
		} else {
			fmt.Printf("receive number is %d \n", curNum)
		}
	}
}

func produce1(channel chan int) {
	for i := 1; i <= 100; i++ {
		channel <- i
	}
	close(channel)
}

func consume1(channel chan int) {
	for {
		curNum, ok := <-channel
		if !ok {
			return
		} else {
			fmt.Printf("receive number is %d \n", curNum)
		}
	}
}
func main() {
	channel := make(chan int)
	go produce(channel)
	go consume(channel)
	time.Sleep(2 * time.Second)
	fmt.Println("---------------------------------")
	channel1 := make(chan int, 100)
	go produce1(channel1)
	go consume1(channel1)
	time.Sleep(2 * time.Second)

}
