package stats

import (
	"sync"
	"testing"
)

func TestRaceCondition(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			IncrementProcessed("jpeg")
		}()
	}
	wg.Wait()
}
