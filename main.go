package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(tasks chan<- int) {
	for i := 1; i <= 8; i++ {
		tasks <- i
	}
	close(tasks)
}

func consumer(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Consumer %d processing task %d\n", id, task)
		time.Sleep(time.Second)
	}
}

func main() {
	tasks := make(chan int) // unbuffered channel
	var wg sync.WaitGroup

	wg.Add(2) // Two consumers

	go producer(tasks)
	go consumer(1, tasks, &wg)
	go consumer(2, tasks, &wg)

	wg.Wait()
}