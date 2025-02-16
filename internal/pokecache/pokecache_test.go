package pokecache

import (
	"testing"
	"time"
)

func TestCreatCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Errorf("Cache is nil")
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	keyone := "key1"
	cache.Add(keyone, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(keyone)
	if ok {
		t.Errorf("%s should be removed", keyone)
	}
}
