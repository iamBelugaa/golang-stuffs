package drop_pattern

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func DropPattern() {
	const work = 2000
	const capacity = 100

	var totalWorkDone int32
	goroutines := runtime.NumCPU()
	dataChan := make(chan int, capacity)

	for i := 0; i < goroutines; i++ {
		go func() {
			for data := range dataChan {
				atomic.AddInt32(&totalWorkDone, 1)
				fmt.Println("Data received", data)
			}
		}()
	}

	for i := 0; i < work; i++ {
		select {
		case dataChan <- work:
		default:
			fmt.Println("DROP")
		}
	}

	close(dataChan)
	fmt.Println("Shutdown signal")
	fmt.Printf("Total work done : %d\n", totalWorkDone)
}
