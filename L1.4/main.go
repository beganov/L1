package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //работа с контекстом позволяет использовать стандартизированный механизм,
			// вместо создания собственного решения. Также контекст обеспечивает безопасность,
			// иерархичность и интеграцию с другими пакетами.
			fmt.Printf("Worker %d: shutting down\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		fmt.Printf("Caught signal: %s\n", sig)
		cancel()
	case <-ctx.Done():
	}
	wg.Wait()
	fmt.Println("All workers shut down.")
}
