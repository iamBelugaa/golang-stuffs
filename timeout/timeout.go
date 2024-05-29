package timeout

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Timeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*150)
	defer cancel()

	dataChan := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		dataChan <- "data"
	}()

	select {
	case <-dataChan:
	case <-ctx.Done():
		fmt.Println("Timeout")
	}

	time.Sleep(time.Second)
	println("Done")
}
