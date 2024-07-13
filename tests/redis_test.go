package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/cache`
)

func TestRedis(t *testing.T) {
	c := cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      7 * 24 * 3600,
	}
	rdx := c.Connection()
	defer rdx.Close()
	err := rdx.Set("account", `{"fullName": "何烨霖", "email": "chaodoing@163.com", "mobile": "15925160015", "password": ""}`, 0).Err()
	if err != nil {
		t.Error(err)
	}
	t.Log(rdx.Get("test").Val())
}
