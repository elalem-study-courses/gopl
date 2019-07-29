package reader

import (
	"sync"
)

type entry struct {
	name       string
	readerFunc func(string) error
}

type registry []*entry

var (
	mu      sync.Mutex
	entries registry
)

func init() {
	entries = make(registry, 0)
}

func RegisterReader(name string, readerFunc func(filePath string) error) {
	mu.Lock()
	defer mu.Unlock()

	entries = append(entries, &entry{name, readerFunc})
}

func TryAllForFile(filePath string) map[string]error {
	errs := map[string]error{}
	for _, entry := range entries {
		err := entry.readerFunc(filePath)
		if err != nil {
			errs[entry.name] = err
		}
	}

	return errs
}
