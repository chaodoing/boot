package container

import (
	`time`
	
	`github.com/chaodoing/boot/auth/captcha`
	`github.com/chaodoing/boot/cache`
	`github.com/go-redis/redis`
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

// Redis 返回 Container 实例中的 Redis 客户端。
// 如果 Redis 客户端尚未初始化，则会根据配置信息进行初始化。
// 这个方法确保了 Redis 客户端的延迟加载和单一实例特性，
// 从而提高资源利用率和性能。
func (c *Container) Redis() *redis.Client {
	// 检查是否已经存在初始化的 Redis 客户端
	if c.rdx != nil {
		return c.rdx
	}
	// 初始化 Redis 客户端并返回
	c.rdx = c.Config.Cache.Connection()
	return c.rdx
}

// Cache 返回缓存实例。如果缓存实例已经存在，则直接返回，否则根据配置信息创建一个新的缓存实例。
func (c *Container) Cache(prefixes ...string) (Cache *cache.Cache, err error) {
	if c.cache != nil && c.rdx != nil {
		return c.cache, nil
	}
	if c.rdx == nil {
		c.rdx = c.Config.Cache.Connection()
	}
	c.cache, err = cache.New(c.rdx, prefixes...)
	_ = c.Events.Trigger(`cache`, c.cache, err)
	return c.cache, err
}

// Group 返回缓存组实例。如果缓存组实例已经存在，则直接返回，否则根据配置信息创建一个新的缓存组实例。
func (c *Container) Group(name ...string) (Group *cache.Group, err error) {
	if c.group != nil && c.rdx != nil {
		return c.group, nil
	}
	if c.rdx == nil {
		c.rdx = c.Config.Cache.Connection()
	}
	c.group, err = cache.NewGroup(c.rdx, name...)
	_ = c.Events.Trigger(`group`, c.group, err)
	return c.group, err
}

// Captcha 方法用于创建并返回一个新的验证码生成器实例。
// 该方法首先通过缓存配置建立到缓存的连接，然后使用验证码配置和缓存连接来创建验证码实例。
// 返回的验证码实例可以用于生成和验证验证码。
//
// 返回值：
// *captcha.Captcha: 返回一个指向 captcha.Captcha 类型的指针。
func (c *Container) Captcha() (cap *captcha.Captcha, err error) {
	// 获取到缓存的连接。
	var rdx = c.Config.Cache.Connection()
	err = rdx.Ping().Err()
	// 使用缓存连接和验证码配置创建并返回一个新的验证码实例。
	cap = captcha.NewCaptcha(c.Config.Captcha, rdx)
	return
}
