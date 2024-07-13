package cache

import (
	`fmt`
	
	`github.com/go-redis/redis`
)

type Config struct {
	Host     string `json:"host" xml:"host" yaml:"host" ini:"HOST" comment:"缓存 REDIS 主机"`
	Port     int    `json:"port" xml:"port" yaml:"port" ini:"PORT" comment:"缓存 REDIS 端口"`
	Password string `json:"password" xml:"password" yaml:"password" ini:"PASSWORD" comment:"连接密码"`
	Index    int    `json:"index" xml:"index" yaml:"index" ini:"INDEX" comment:"缓存索引"`
	TTL      int    `json:"ttl" xml:"ttl" yaml:"ttl" ini:"TTL" comment:"缓存默认过期时间"`
}

func (c *Config) Connection() (rdx *redis.Client) {
	rdx = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password,
		DB:       c.Index,
	})
	return
}
