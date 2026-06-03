package processor

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

var (
	LeakCache    = make(map[string][]byte)
	leakCacheMu  sync.Mutex
	maxCacheSize = 1000

	// Оптимізація CPU: виносимо компіляцію
	imageRe = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)
)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	select {}
}

func processImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	// Використовуємо вже скомпільований вираз
	if imageRe.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())

		leakCacheMu.Lock()
		if len(LeakCache) >= maxCacheSize {
			LeakCache = make(map[string][]byte, maxCacheSize)
		}
		LeakCache[key] = make([]byte, 1024*10)
		leakCacheMu.Unlock()
	}
}
