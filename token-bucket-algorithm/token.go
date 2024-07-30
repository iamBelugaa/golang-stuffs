package token_bucket_algorithm

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	sync.Mutex
	Tokens     int
	Capacity   int
	StopRefill chan struct{}
	RefillRate time.Duration
}

func NewTokenBucket(capacity, tokensPerInterval int, refillRate time.Duration) *TokenBucket {
	tb := TokenBucket{
		Capacity:   capacity,
		RefillRate: refillRate,
		StopRefill: make(chan struct{}),
	}

	go tb.refillTokens(tokensPerInterval)
	return &tb
}

func (tb *TokenBucket) refillTokens(tokensPerInterval int) {
	ticker := time.NewTicker(tb.RefillRate)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tb.Lock()

			if tb.Tokens+tokensPerInterval <= tb.Capacity {
				tb.Tokens += tokensPerInterval
			} else {
				tb.Tokens = tb.Capacity
			}

			tb.Unlock()
		case <-tb.StopRefill:
			return
		}
	}
}

func (tb *TokenBucket) TakeTokens() bool {
	tb.Lock()
	defer tb.Unlock()

	if tb.Tokens > 0 {
		tb.Tokens--
		return true
	}

	return false
}

func (tb *TokenBucket) StopRefillFunc() {
	close(tb.StopRefill)
}

func Run() {
	tb := NewTokenBucket(10, 5, time.Second)

	time.Sleep(time.Second)

	for i := 1; i <= 100; i++ {
		if tb.TakeTokens() {
			fmt.Printf("Token taken. Remaining tokens: %d\n", tb.Tokens)
		} else {
			fmt.Printf("Not enough tokens.\n")
			time.Sleep(350 * time.Millisecond)
		}
	}

	tb.StopRefillFunc()
}
