package tests

import (
	`fmt`
	`testing`
	`time`
	
	`github.com/chaodoing/boot/cache`
)

func TestCache(t *testing.T) {
	config1 := &cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      7 * 24 * 60 * 60,
	}
	rdx := config1.Connection()
	rdx.Set("test_key", "test_value", time.Duration(3600)*time.Second)
	rdx.Set("account", `{"username": "chaodoing", "email": "chaodoing@live.com", "mobile": "15925160015", "gender": 0}`, time.Duration(3600)*time.Second)
	t.Log(rdx.Get("test_key"))
	t.Log(rdx.Get("account"))
	config := &cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      7 * 24 * 60 * 60,
	}
	Cache, err := cache.New(config)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		var key = fmt.Sprintf("test_key_%d", i)
		var value = fmt.Sprintf("test_value_%d", i)
		Cache.Set(key, value, 3600)
		t.Log(Cache.Get(key))
	}
	if err = Cache.Clear(); err != nil {
		t.Error(err)
	}
}
