package main

import (
	"fmt"
	"sync"
)

// L1.9
// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout.
// То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

var array []int = []int{1, 3, 8, 8, 6, 0, 4}

func main() {
	// Даны два канала
	firstChan := make(chan int)
	secondChan := make(chan int)
	var wg sync.WaitGroup

	// в первый пишутся числа x из массива
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(firstChan)
		for _, i := range array {
			firstChan <- i
		}
	}()

	// во второй – результат операции x*2.
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(secondChan)
		for data := range firstChan {
			secondChan <- data * 2
		}
	}()

	// После этого данные из второго канала должны выводиться в stdout.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range secondChan {
			fmt.Println(data)
		}
	}()
	wg.Wait()
}
