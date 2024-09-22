package container

import (
	`github.com/chaodoing/boot/captcha`
)

// Captcha 方法用于创建并返回一个新的验证码生成器实例。
// 该方法首先通过缓存配置建立到缓存的连接，然后使用验证码配置和缓存连接来创建验证码实例。
// 返回的验证码实例可以用于生成和验证验证码。
//
// 返回值：
// *captcha.Captcha: 返回一个指向 captcha.Captcha 类型的指针。
func (c *Container) Captcha() *captcha.Captcha {
	// 获取到缓存的连接。
	var rdx = c.Config.Cache.Connection()
	// 使用缓存连接和验证码配置创建并返回一个新的验证码实例。
	return captcha.NewCaptcha(c.Config.Captcha, rdx, 5)
}
