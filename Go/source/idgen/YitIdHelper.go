package idgen

import (
	"sync"
)

var singletonMutex sync.Mutex
var idGenerator *DefaultIdGenerator

// SetIdGenerator .
func SetIdGenerator(options *IdGeneratorOptions) {
	singletonMutex.Lock()
	idGenerator = NewDefaultIdGenerator(options)
	singletonMutex.Unlock()
}

// NextId .
func NextId() uint64 {
	if idGenerator == nil {
		singletonMutex.Lock()
		defer singletonMutex.Unlock()
		if idGenerator == nil {
			options := NewIdGeneratorOptions(1)
			idGenerator = NewDefaultIdGenerator(options)
		}
	}

	return idGenerator.NewLong()
}
