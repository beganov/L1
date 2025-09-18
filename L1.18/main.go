package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// L1.18
// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде
// (т.е. из нескольких горутин). По завершению программы структура должна выводить итоговое значение счётчика.
// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.

type Counter struct { // Структура-счётчик
	counter atomic.Int64 //sync/Atomic для безопасного инкремента
}

func New() *Counter {
	var Counter Counter
	Counter.counter.Store(0)
	return &Counter
}

func main() {
	Counter := New()
	var wg sync.WaitGroup
	unsafeCounter := 0 //потоко небезопасный счетчик, для сравнения
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeCounter++
			Counter.counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(Counter.counter.Load(), unsafeCounter)
	// unsafeCounter стабильно теряет процент-два, атомик - нет
}
