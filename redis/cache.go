package cache

import (
	`github.com/go-redis/redis`
)

type Cache struct {
	rdx   *redis.Client
	cache map[string]*redis.Client
}

func (c Cache) Get(p, key string) {

}

func (c Cache) Set(p, key string) {

}
