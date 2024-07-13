package cache

import (
	`errors`
	`time`
	
	`github.com/go-redis/redis`
)

// Cache 是一个基于 Redis 的缓存结构体，包含 Redis 配置、Redis 客户端和键前缀。
type Cache struct {
	redisConfig *Config
	rdx         *redis.Client
	prefix      string
}

// redis 用于初始化 Redis 连接，并检查连接是否成功。
func (c *Cache) redis() error {
	c.rdx = c.redisConfig.Connection()
	if c.rdx == nil {
		return errors.New("redis connection error")
	}
	return c.rdx.Ping().Err()
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
