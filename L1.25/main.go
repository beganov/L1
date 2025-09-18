package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// L1.25
// Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep,
// которая приостанавливает выполнение текущей горутины.
// Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение
// (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.
// Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).

var duration = 2 * time.Second

func main() {
	var wg sync.WaitGroup
	fmt.Println("Start")
	now := time.Now()

	wg.Add(1) //через контекст
	go func() {
		defer wg.Done()
		sleepCtx(duration)
		fmt.Println("Wake up 1", time.Since(now))
	}()

	wg.Add(1) //через цикл
	go func() {
		defer wg.Done()
		sleepFor(duration)
		fmt.Println("Wake up 2", time.Since(now))
	}()

	wg.Add(1) //через timeAfter
	go func() {
		defer wg.Done()
		sleepTimeAfter(duration)
		fmt.Println("Wake up 3", time.Since(now))
	}()

	wg.Add(1) //через timer
	go func() {
		defer wg.Done()
		sleepTimer(duration)
		fmt.Println("Wake up 4", time.Since(now))
	}()

	wg.Wait()
}

func sleepCtx(duration time.Duration) { //ждем, пока жив контест
	ctx, canel := context.WithTimeout(context.Background(), duration)
	defer canel()
	<-ctx.Done()
}

func sleepFor(duration time.Duration) { //ждем, пока не выйдем из цикла
	now := time.Now()
	for {
		if time.Since(now) > duration {
			return
		}
	}
}

func sleepTimeAfter(duration time.Duration) { //блокировка через time.After
	<-time.After(duration)
}

func sleepTimer(duration time.Duration) { //блокировка через timer
	timer := time.NewTimer(duration)
	<-timer.C
}
