package container

import (
	`time`
	
	`github.com/chaodoing/boot/cache`
	`gorm.io/gorm`
)

// Jwt 返回一个Jwt实例。如果实例已经存在，则直接返回，否则根据配置信息创建一个新的Jwt实例。
func (c *Container) Jwt() *Jwt {
	if c.jwt != nil {
		return c.jwt
	}
	c.jwt = NewJwt(c.Config.Jwt.Secret, time.Duration(c.Config.Jwt.Expire)*time.Second)
	_ = c.Events.Trigger(`jwt`, c.jwt)
	return c.jwt
}

// Database 返回数据库连接。如果数据库连接已经建立，则直接返回，否则根据配置信息建立新的数据库连接。
func (c *Container) Database() (*gorm.DB, error) {
	if c.db != nil {
		return c.db, nil
	}
	var err error
	c.db, err = c.Config.Database.Connection()
	_ = c.Events.Trigger(`database`, c.db, err)
	return c.db, err
}

// Cache 返回缓存实例。如果缓存实例已经存在，则直接返回，否则根据配置信息创建一个新的缓存实例。
func (c *Container) Cache(prefixes ...string) (Cache *cache.Cache, err error) {
	if c.cache != nil {
		return c.cache, nil
	}
	c.cache, err = cache.New(&c.Config.Cache, prefixes...)
	_ = c.Events.Trigger(`cache`, c.cache, err)
	return c.cache, err
}

// Group 返回缓存组实例。如果缓存组实例已经存在，则直接返回，否则根据配置信息创建一个新的缓存组实例。
func (c *Container) Group(name ...string) (Group *cache.Group, err error) {
	if c.group != nil {
		return c.group, nil
	}
	c.group, err = cache.NewGroup(&c.Config.Cache, name...)
	_ = c.Events.Trigger(`group`, c.group, err)
	return c.group, err
}
