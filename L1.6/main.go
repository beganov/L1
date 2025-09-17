package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления, через контекст,
// прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

func main() {
	stopFlag := false
	stopChan := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go func(stopFlag *bool) { //по условию
		defer wg.Done()
		for !*stopFlag {
			fmt.Println("Worker 1: working...")
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Worker 1 stoped.")
	}(&stopFlag)

	wg.Add(1)
	go func(chan struct{}) { //через канал
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				fmt.Println("Worker 2 stoped.")
				return
			default:
				fmt.Println("Worker 2: working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(stopChan)

	wg.Add(1)
	go func(ctx context.Context) { //через контекст
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker 3 stoped.")
				return
			default:
				fmt.Println("Worker 3: working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)

	wg.Add(1)
	go func() { //через runtime.Goexit()
		defer wg.Done()
		counter := 0
		for {
			fmt.Println("Worker 4: working...")
			time.Sleep(1 * time.Second)
			counter++
			if counter >= 3 {
				fmt.Println("Worker 4 stoped.")
				runtime.Goexit()
			}
		}
	}()

	wg.Add(1)
	go func() { // Через контролируемую panic, явно не классический, но горутину остановит
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Worker 5 stopped.", r)
			}
			wg.Done()
		}()
		counter := 0
		for {
			fmt.Println("Worker 5: working...")
			time.Sleep(1 * time.Second)
			counter++
			if counter >= 3 {
				panic("panic")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	stopFlag = true
	close(stopChan)
	cancel()
	wg.Wait()
}
