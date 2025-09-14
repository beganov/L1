package main

import (
	"fmt"
	"sync"
)

var array []int = []int{1, 3, 8, 8, 6, 0, 4}

func main() {
	firstChan := make(chan int)
	secondChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(firstChan)
		for _, i := range array {
			firstChan <- i
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(secondChan)
		for data := range firstChan {
			secondChan <- data * 2
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range secondChan {
			fmt.Println(data)
		}
	}()
	wg.Wait()
}
