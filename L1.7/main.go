package main

import (
	"fmt"
	"sync"
)

type Cache struct {
	sync.Mutex
	mapInt map[int]int
}

func main() {
	che := New()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			che.set(i)
		}(i)
	}
	wg.Wait()
}

func New() *Cache {
	return &Cache{mapInt: make(map[int]int, 10)}
}

func (c *Cache) set(i int) {
	defer c.Unlock()
	c.Lock()
	c.mapInt[i]++
	fmt.Println(i, c.mapInt[i])
}
