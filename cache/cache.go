package cache

import (
	`fmt`
	`strings`
	`time`
	
	`github.com/go-redis/redis`
)

// Cache 是一个基于 Redis 的缓存结构体，包含 Redis 配置、Redis 客户端和键前缀。
type Cache struct {
	rdx    *redis.Client
	prefix string
}

// All 返回当前分组下所有键值对。
// 它通过匹配前缀来获取键值，并移除键的前缀以得到最终的结果。
func (c *Cache) All() map[string]string {
	var result = map[string]string{}
	names := c.rdx.Keys(fmt.Sprintf("%s:*", c.prefix)).Val()
	key := fmt.Sprintf("%s:", c.prefix)
	for _, name := range names {
		result[strings.TrimPrefix(name, key)] = c.rdx.Get(name).Val()
	}
	return result
}

// Set 用于设置缓存值。key 为缓存键，value 为缓存值，ttl 为可选的缓存过期时间（秒）。
func (c *Cache) Set(key string, value interface{}, ttl ...int) error {
	var duration = time.Duration(0)
	if len(ttl) > 0 {
		duration = time.Duration(ttl[0]) * time.Second
	}
	key = c.prefix + ":" + key
	return c.rdx.Set(key, value, duration).Err()
}

// Get 用于获取缓存值。key 为缓存键。
func (c *Cache) Get(key string) string {
	key = c.prefix + ":" + key
	return c.rdx.Get(key).Val()
}

// Exist 用于检查缓存键是否存在。
func (c *Cache) Exist(key string) bool {
	key = c.prefix + ":" + key
	return c.rdx.Exists(key).Val() > 0
}

// Expire 用于获取缓存键的过期时间。
func (c *Cache) Expire(key string) time.Duration {
	key = c.prefix + ":" + key
	return c.rdx.TTL(key).Val()
}

// Delete 用于删除缓存键。
func (c *Cache) Delete(key string) error {
	key = c.prefix + ":" + key
	return c.rdx.Del(key).Err()
}

// Clear 用于清除指定前缀的所有缓存键。
func (c *Cache) Clear() error {
	var key = c.prefix + ":*"
	keys := c.rdx.Keys(key).Val()
	if err := c.rdx.Del(keys...).Err(); err != nil {
		return err
	}
	return nil
}
