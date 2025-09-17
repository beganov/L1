package main

import (
	"fmt"
	"os"
	"strconv"
)

// L1.3
// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого
// канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и
// при старте создавать указанное число горутин-воркеров.

func main() {
	if len(os.Args) < 2 { //если не передано количество воркеров
		fmt.Println("Usage: go run main.go <number_of_workers>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 { //если количество передано не валидное
		fmt.Println("Please provide a valid number of workers")
		return
	}

	ch := make(chan int)

	//запуск воркеров
	for i := 0; i < n; i++ {
		go worker(i, ch)
	}

	//постоянная запись в канал
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
