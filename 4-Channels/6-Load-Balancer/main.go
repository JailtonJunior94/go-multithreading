package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	quantidadeWorkers := 50

	for i := 0; i < quantidadeWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 100; i++ {
		data <- i
	}
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}
