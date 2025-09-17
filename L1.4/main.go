package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var numWorkers = 3

// L1.4
// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров
// при получении сигнала прерывания.
// Подсказка: можно использовать контекст (context.Context)
// или канал для оповещения о завершении.

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done(): //остановка воркеров при завершении контекста

			// работа с контекстом позволяет использовать стандартизированный механизм,
			// вместо создания собственных решений. Также контекст обеспечивает
			// безопасность, иерархичность и интеграцию с другими пакетами.
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

	// запуск воркеров
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	// настройка канала на регистрацию SIGINT и SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание либо сигнала, либо отмены контекста
	select {
	case sig := <-sigChan:
		fmt.Printf("Caught signal: %s\n", sig)
		cancel()
	case <-ctx.Done():
	}
	wg.Wait()
	fmt.Println("All workers shut down.")
}
