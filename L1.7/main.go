package main

import (
	"fmt"
	"sync"
)

// L1.7
// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

var lenArr = 10

type Cache struct {
	sync.Mutex // реализация через sync.Mutex
	mapInt     map[int]int
}

func main() {
	cache := New()
	var wg sync.WaitGroup
	for i := 0; i < lenArr; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cache.set(i)
		}(i)
	}
	wg.Wait()
}

func New() *Cache {
	return &Cache{mapInt: make(map[int]int, lenArr)}
}

func (c *Cache) set(i int) {
	defer c.Unlock() // открываем при выходе
	c.Lock()         // закрываем мьютекс при входе - иначе будут гонки
	c.mapInt[i]++
	fmt.Println(i, c.mapInt[i])
}

// Проверьте работу кода на гонки (util go run -race).
//go run -race L1.7/main.go
// 0 1
// 1 1
// 2 1
// 9 1
// 5 1
// 6 1
// 3 1
// 8 1
// 4 1
// 7 1

// без Lock/Unlock
// go run -race L1.7/main.go
// ==================
// WARNING: DATA RACE
// Read at 0x00c00008a150 by goroutine 8:
// runtime.mapaccess1_fast64()
//  C:/Program Files/Go/src/runtime/map_fast64.go:13 +0x0
// main.(*Cache).set()
//   C:/Users/began/L1/L1.7/main.go:38 +0x6b
//  main.main.func1()
//   C:/Users/began/L1/L1.7/main.go:25 +0x8b
//  main.main.gowrap1()
//      C:/Users/began/L1/L1.7/main.go:26 +0x41
// Previous write at 0x00c00008a150 by goroutine 7:
//   runtime.mapassign_fast64()
//      C:/Program Files/Go/src/runtime/map_fast64.go:113 +0x0
//   main.(*Cache).set()
//       C:/Users/began/L1/L1.7/main.go:38 +0x9b
//   main.main.func1()
//       C:/Users/began/L1/L1.7/main.go:25 +0x8b
//   main.main.gowrap1()
//      C:/Users/began/L1/L1.7/main.go:26 +0x41
// Goroutine 8 (running) created at:
//   main.main()
//       C:/Users/began/L1/L1.7/main.go:23 +0x117
// Goroutine 7 (running) created at:
//   main.main()
//      C:/Users/began/L1/L1.7/main.go:23 +0x117
// ==================
// ==================
// WARNING: DATA RACE
// ... как будто нет смысла весь лог для подтверждения факта гонки приводить
// 4 1
// 9 1
// Found 3 data race(s)
// exit status 66
