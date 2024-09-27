package token

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// mem cache
var c *cache.Cache

func init() {
	c = cache.New(5*time.Minute, 10*time.Minute)
}

func Set(key string, value interface{}, expiration time.Duration) {
	c.Set(key, value, expiration)
}

func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

func Delete(key string) {
	c.Delete(key)
}
