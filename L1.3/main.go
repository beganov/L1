package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 {
		fmt.Println("Please provide a valid number of workers")
		return
	}

	ch := make(chan int)

	for i := 0; i < n; i++ {
		go worker(i, ch)
	}

	counter := 0
	for {
		ch <- counter
		counter++
	}
}

func worker(id int, ch <-chan int) {
	for data := range ch {
		fmt.Printf("Worker %d: %d\n", id, data)
	}
}
