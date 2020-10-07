package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.RWMutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string, wg *sync.WaitGroup) {
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()

	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()

}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.RLock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.RUnlock()
	return c.v[key]
}

func main() {
	var wg sync.WaitGroup

	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go c.Inc("somekey", &wg)
		go c.Value("somekey")
	}

	wg.Wait()
	fmt.Println(c.Value("somekey"))
}
