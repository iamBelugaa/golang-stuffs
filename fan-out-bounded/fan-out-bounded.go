package fan_out_bounded

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func FanOutBound() {
	var wg sync.WaitGroup

	work := 2_00_00_000
	numCPU := runtime.NumCPU()
	dataChan := make(chan int, numCPU)

	println("Num CPU :", numCPU)

	wg.Add(numCPU)
	start := time.Now()

	for i := 0; i < numCPU; i++ {
		go func() {
			defer wg.Done()
			for data := range dataChan {
				data++
			}
		}()
	}

	for i := 0; i < work; i++ {
		dataChan <- i
	}

	close(dataChan)
	wg.Wait()

	fmt.Printf("\nTime taken for %d work : %dns\n", work, time.Since(start))
}
