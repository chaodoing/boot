package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/cache`
)

func TestRedis(t *testing.T) {
	// c := cache.Config{
	// 	Host:     "r-2vc0zzpisbe452355dpd.redis.cn-chengdu.rds.aliyuncs.com",
	// 	Port:     6379,
	// 	Password: "heyelin:hyl9123456789852Hyl",
	// 	Index:    0,
	// 	TTL:      7 * 24 * 3600,
	// }
	c1 := cache.Config{
		Host:     "192.168.33.10",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      0,
	}
	// rdx := c.Connection()
	rdx1 := c1.Connection()
	defer func() {
		// _ = rdx.Close()
		_ = rdx1.Close()
	}()
	var data, err = rdx1.LRange("luxiaoqian:order:449:locus", 0, -1).Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(data)
	
	// var data, err = rdx.LRange("luxiaoqian:order:449:locus", 0, -1).Result()
	// if err != nil {
	// 	t.Error(err)
	// }
	// for _, datum := range data {
	// 	t.Log(datum)
	// }
}
