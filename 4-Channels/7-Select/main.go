package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- 1
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("received", msg1)
	case msg2 := <-c2:
		fmt.Println("received", msg2)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout")
	default:
		fmt.Println("default")
	}
}
