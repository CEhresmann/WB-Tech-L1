package main

import (
	"fmt"
	"sync"
	"time"
) //the map type in Go doesn't support concurrent reads and writes.

// concurrent-map provides a high-performance solution to this by sharding the map with minimal time spent waiting for locks.
//это относится к устаревшей библиотеке cmap "github.com/orcaman/concurrent-map/v2"
//но я использовал Map из библиотеки sync
// использование библиотеки упрощает взаимодействие с отображениями, для которых необходимы конкурентная запись и чтение.
// Но в дальнейшем коде реализованы также оригинальные функции

type ConcurrentMap struct {
	mu    sync.RWMutex
	store map[string]int
}

// Set Метод для записи данных в отображение
func (c *ConcurrentMap) Set(key string, value int) {
	c.mu.Lock()         // Захватываем мьютекс для записи
	defer c.mu.Unlock() // Освобождаем мьютекс после завершения функции
	c.store[key] = value
}

// Get Метод для чтения данных из отображения
func (c *ConcurrentMap) Get(key string) (int, bool) {
	c.mu.RLock()         // Захватываем мьютекс для чтения
	defer c.mu.RUnlock() // Освобождаем мьютекс после завершения функции
	value, ok := c.store[key]
	return value, ok
}

func main() {
	var m sync.Map
	cMap := &ConcurrentMap{
		store: make(map[string]int),
	}
	var wg sync.WaitGroup

	// Запуск горутин для записи в отображение
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			cMap.Set(fmt.Sprintf("key %d", i), i)
		}(i)
	}

	// Запуск горутин для чтения из отображения
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond) // or 100000 Nanoseconds (empirically)
			if value, ok := cMap.Get(fmt.Sprintf("key %d", i)); ok {
				fmt.Printf("key %d: %d\n", i, value)
			} else {
				fmt.Printf("key %d not found\n", i)
			}
		}(i)
	}

	// Запуск горутин для записи в sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(fmt.Sprintf("key %d", i), i)
			time.Sleep(time.Millisecond)
		}(i)
	}

	// Запуск горутин для чтения из sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			if value, ok := m.Load(fmt.Sprintf("key %d", i)); ok {
				fmt.Printf("key %d of sync.Map: %d\n", i, value)
			} else {
				fmt.Printf("key %d of sync.Map  not found\n", i)
			}
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин
}
