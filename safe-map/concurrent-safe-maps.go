package safe_maps

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Shard struct {
	mu   sync.RWMutex
	data map[string]any
}

type ShardMap []*Shard

func NewShardMap(cap int) ShardMap {
	shards := make(ShardMap, cap)

	for i := 0; i < cap; i++ {
		shards[i] = &Shard{
			data: make(map[string]any),
		}
	}

	return shards
}

func (m ShardMap) getShardIndex(key string) int {
	checksum := sha256.Sum256([]byte(key))
	return int(checksum[0]) % len(m)
}

func (m ShardMap) getShard(key string) *Shard {
	i := m.getShardIndex(key)
	fmt.Printf("Key %s and Shard %d \n", key, i)
	return m[i]
}

func (m ShardMap) Get(key string) (any, bool) {
	shard := m.getShard(key)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	val := shard.data[key]
	return val, val != nil
}

func (m ShardMap) Set(key string, val any) {
	shard := m.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	shard.data[key] = val
}

func (m ShardMap) Delete(key string) {
	shard := m.getShard(key)

	shard.mu.Lock()
	defer shard.mu.Unlock()

	delete(shard.data, key)
}

func (m ShardMap) Contains(key string) bool {
	shard := m.getShard(key)

	shard.mu.RLock()
	defer shard.mu.RUnlock()

	v := shard.data[key]
	return v != nil
}

func (m ShardMap) Keys() []string {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	keys := make([]string, 0)

	wg.Add(len(m))

	for _, shard := range m {
		go func(s *Shard) {
			s.mu.RLock()
			defer wg.Done()
			defer s.mu.RUnlock()

			for key := range s.data {
				mutex.Lock()
				keys = append(keys, key)
				mutex.Unlock()
			}
		}(shard)
	}

	wg.Wait()
	return keys
}

func ConcurrentSafeMaps() {
	cache := NewShardMap(10)

	cache.Set("abcd", "48")
	cache.Set("hello", "world")

	cache.Delete("abcd")
	cache.Delete("hello")

	fmt.Printf("\nKeys : %+v\n", cache.Keys())
	fmt.Printf("1 exists : %v\n", cache.Contains("1"))
	fmt.Printf("2 exists : %v\n", cache.Contains("2"))
	fmt.Printf("abcd exists : %v\n", cache.Contains("abcd"))
	fmt.Printf("Hello exists : %v\n", cache.Contains("hello"))
}
