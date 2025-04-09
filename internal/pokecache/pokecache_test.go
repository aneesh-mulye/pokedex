package pokecache_test

import (
	"github.com/aneesh-mulye/pokedex/internal/pokecache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	a := assert.New(t)
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache := pokecache.NewCache(interval)
		cache.Add(c.key, c.val)
		val, ok := cache.Get(c.key)
		a.True(ok, "expected to find key %s", c.key)
		a.Equal(val, c.val, "expected to find value %s", string(c.val))
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
