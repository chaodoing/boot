package tests

import (
	`encoding/json`
	`fmt`
	`testing`
	
	`github.com/chaodoing/boot/cache`
)

func TestGroup(t *testing.T) {
	config := cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      0,
	}
	rdx := config.Connection()
	g, err := cache.NewGroup(rdx, "cache", "group")
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		var key = fmt.Sprintf("test_key_%d", i)
		var value = fmt.Sprintf("test_value_%d", i)
		g.Set(key, value, 3600)
	}
	j, err := json.MarshalIndent(g.All(), "", "\t")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(j))
	if err = g.Clear(); err != nil {
		t.Error(err)
	}
}
