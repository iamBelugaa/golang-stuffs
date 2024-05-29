package drop_pattern

import (
	"fmt"
	"runtime"
)

func DropPattern() {
	const work = 2000
	const capacity = 100

	goroutines := runtime.NumCPU()
	dataChan := make(chan int, capacity)

	for i := 0; i < goroutines; i++ {
		go func() {
			for data := range dataChan {
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
}
