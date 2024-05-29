package confinement

import (
	"fmt"
	"sync"
)

func processData(wg *sync.WaitGroup, result []int, index, data int) {
	defer wg.Done()
	result[index] = data * 2
}

func Confinement() {
	var wg sync.WaitGroup
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, len(input))

	wg.Add(len(input))

	for i, data := range input {
		go processData(&wg, result, i, data)
	}

	wg.Wait()
	fmt.Printf("Result : %+v", result)
}
