package container

import (
	`github.com/chaodoing/boot/cache`
	`github.com/chaodoing/boot/config`
	`github.com/chaodoing/boot/task`
	`github.com/go-redis/redis`
	`gorm.io/gorm`
)

type Container struct {
	Events  *task.Events
	Crontab *task.Crontab
	Config  config.Config
	
	cache *cache.Cache
	group *cache.Group
	jwt   *Jwt
	db    *gorm.DB
	rdx   *redis.Client
}
