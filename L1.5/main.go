package main

import (
	"context"
	"fmt"
	"time"
)

const (
	n = 4
)

func worker(ch <-chan int, ctx context.Context, timeout <-chan time.Time) {
	for data := range ch {
		select {
		//case <-ctx.Done():
		case <-timeout:
			fmt.Printf("worker shutting down\n")
			return
		default:
			fmt.Printf("Data: %d\n", data)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	//как будто бы логично добавить таймаут в контекст
	defer cancel()

	ch := make(chan int)
	timer := time.Now()
	timeout := time.After(n * time.Second)
	//но в подсказке time.After, так что реализация с ним
	go worker(ch, ctx, timeout)
	counter := 0
	for {
		select {
		//case <-ctx.Done():
		case <-timeout:
			fmt.Println("context is done", time.Since(timer))
			close(ch)
			return
		case ch <- counter:
			counter++
		}
	}
}
