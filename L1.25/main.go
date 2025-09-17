package main

import (
	"context"
	"fmt"
	"time"
)

var duration = 2 * time.Second

func main() {
	fmt.Println("Start")
	now := time.Now()
	sleepCtx(duration)
	fmt.Println("Wake up 1", time.Since(now))
	now = time.Now()
	sleepFor(duration)
	fmt.Println("Wake up 2", time.Since(now))
	now = time.Now()
	sleepTimeAfter(duration)
	fmt.Println("Wake up 3", time.Since(now))
}

func sleepCtx(duration time.Duration) {
	ctx, canel := context.WithTimeout(context.Background(), duration)
	defer canel()
	<-ctx.Done()
}

func sleepFor(duration time.Duration) {
	now := time.Now()
	for {
		if time.Since(now) > duration {
			return
		}
	}
}

func sleepTimeAfter(duration time.Duration) {
	<-time.After(duration)
}
