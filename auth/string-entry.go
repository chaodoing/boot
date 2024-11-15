package auth

import (
	`time`
	
	`github.com/go-redis/redis`
)

const (
	Basic  = "Basic "
	Bearer = "Bearer "
)

var HeaderKeys = []string{
	"Authorization",
	"Access-Token",
	"Token",
	"X-Websocket-Header-Authorization",
	"X-Websocket-Header-Token",
	"X-Websocket-Header-Access-Token",
}

// StringEntry 通过生成的随机数作为用户认证字符串
type StringEntry struct {
	rdx        *redis.Client
	timeToLive time.Duration
	prefix     string
}
