package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	Instance *cache.Cache
	once     sync.Once
)

// M returns a singleton.
func M() *cache.Cache {
	Init()
	return Instance
}

// Init inits the singleton.
func Init() {
	once.Do(func() {
		Instance = cache.New(5*time.Minute, 10*time.Minute)
	})
}
