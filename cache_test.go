package monnify

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	cache.Set("test", "testing", 2*time.Second)

	val, _ := cache.Get("test")
	assert.Equal(t, "testing", val, "Should return 'testing'")

	time.Sleep(2 * time.Second)
	val, _ = cache.Get("test")
	assert.Equal(t, "", val, "value should be empty")

	cache.Set("test", "testing", 2*time.Second)
	cache.Delete("test")
}
