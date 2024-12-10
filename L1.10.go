package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const numTemperatures = 50
	temperatures := make([]float64, numTemperatures)
	for i := 0; i < numTemperatures; i++ {
		temperatures[i] = rand.Float64()*70 - 30 // 71 = 40 - (-30) + 1
	}

	TEMPS := make(map[int][]float64)

	for _, temp := range temperatures {
		groupKey := int(temp/10) * 10
		TEMPS[groupKey] = append(TEMPS[groupKey], temp)
	}

	keys := make([]int, 0, len(TEMPS))
	for key := range TEMPS {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Printf("Сгенерированные температуры: %.1f\n", temperatures)
	fmt.Println("Группы температур:")
	for _, values := range keys {
		fmt.Printf("%d: %.1f\n", values, TEMPS[values])
	}
}
