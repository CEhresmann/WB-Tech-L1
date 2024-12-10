package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
)

type Count struct {
	value int64
	mu    sync.Mutex
}

func (c *Count) Inc1() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}
func (c *Count) Inc2() {
	atomic.AddInt64(&c.value, 1)
}

func (c *Count) Value1() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *Count) Value2() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	ct := &Count{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				ct.Inc1()
				if ct.value%33 == 0 && rand.Intn(1000) <= 5 {
					fmt.Println("Значение инкрементации первым типом", ct.Value1())
				}
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				ct.Inc2()
				if ct.value%33 == 0 && rand.Intn(1000) <= 5 {
					fmt.Println("Значение инкрементации вторым(atomic) типом", ct.Value1())
				}
			}
		}()
	}
	wg.Wait()
	fmt.Printf("\n\n\t\t\t\tИтоговое значение после работы: %d\n", ct.Value1())
}
