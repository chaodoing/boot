package tests

import (
	`fmt`
	`testing`
	
	`github.com/chaodoing/boot/cache`
)

func TestCache(t *testing.T) {
	config := &cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      7 * 24 * 60 * 60,
	}
	rdx := config.Connection()
	Cache, err := cache.New(rdx)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		var key = fmt.Sprintf("test_key_%d", i)
		var value = fmt.Sprintf("test_value_%d", i)
		Cache.Set(key, value, 3600)
		// t.Log(Cache.Get(key))
	}
	t.Log(Cache.All())
	if err = Cache.Clear(); err != nil {
		t.Error(err)
	}
}
