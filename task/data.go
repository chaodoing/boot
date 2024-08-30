package task

import (
	`github.com/chaodoing/boot/cache`
	`github.com/chaodoing/boot/container`
	`gorm.io/gorm`
)

var Event = map[string]interface{}{
	"jwt":      func(jwt *container.Jwt) {},
	"database": func(db *gorm.DB, err error) {},
	"cache":    func(cache *cache.Cache, err error) {},
	"group":    func(group *cache.Group, err error) {},
}
