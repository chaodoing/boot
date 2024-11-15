package captcha

import (
	`github.com/go-redis/redis`
	`github.com/mojocn/base64Captcha`
)

type (
	Options struct {
		Height   int     `json:"height" xml:"height" ini:"HEIGHT" comment:"图片高度"`
		Width    int     `json:"width" xml:"width" ini:"WIDTH" comment:"图片宽度"`
		Length   int     `json:"length" xml:"length" ini:"LENGTH" comment:"内容长度"`
		MaxSkew  float64 `json:"max_skew" xml:"max_skew" ini:"MAX_SKEW" comment:"文字倾斜角度"`
		DotCount int     `json:"dot_count" xml:"dot_count" ini:"DOT_COUNT" comment:"背景杂色点数量"`
		InDate   int64   `json:"indate" xml:"indate" ini:"INDATE" comment:"验证码有效期单位:分钟"`
	}
	Captcha struct {
		driver  *base64Captcha.DriverDigit
		store   Store
		captcha *base64Captcha.Captcha
	}
)

// NewCaptcha 创建并返回一个新的验证码实例。
//
// 该函数接收配置选项、Redis客户端和有效期作为参数，用于初始化验证码的驱动、存储层和验证码本身。
// 主要用于系统中需要生成验证码的地方，通过Redis来存储验证码的数据，以验证用户输入的验证码是否正确。
//
// 参数:
//   option - 验证码的配置选项，包括尺寸、长度等参数。
//   rdx - Redis客户端，用于存储验证码数据。
//   indate - 验证码的有效期。
//
// 返回值:
//   *Captcha - 一个指向初始化后的验证码实例的指针。
func NewCaptcha(option Options, rdx *redis.Client) *Captcha {
	// 创建一个新的验证码驱动实例，用于生成数字验证码图片。
	var driver = base64Captcha.NewDriverDigit(option.Height, option.Width, option.Length, option.MaxSkew, option.DotCount)
	// 创建一个新的验证码存储实例，使用Redis作为存储后端。
	var store = NewStore(rdx, option.InDate)
	
	// 返回一个新的验证码实例，使用之前创建的驱动和存储层进行初始化。
	return &Captcha{
		driver:  driver,
		store:   store,
		captcha: base64Captcha.NewCaptcha(driver, base64Captcha.Store(store)),
	}
}

func (c *Captcha) Clear(key string) error {
	return c.store.Clear(key)
}

// Base64Image 生成并返回验证码的Base64编码字符串。
// 这个方法调用captcha结构的Generate方法来生成验证码。
// 返回的id是验证码的唯一标识符，用于后续的验证过程。
// 返回的image是Base64编码后的验证码图片数据，可以直接传输。
// 返回的answer是验证码的正确答案，需要在验证时与用户输入进行对比。
// 如果生成过程中发生错误，err将返回非nil值。
func (c *Captcha) Base64Image() (id, image, answer string, err error) {
	return base64Captcha.NewCaptcha(c.driver, c.store).Generate()
}

// Verify 验证码验证函数
// 通过调用存储层的 Verify 方法来校验验证码的正确性
// 参数:
//   idKey: 验证码的唯一标识符，用于在存储中查找验证码
//   value: 用户提交的验证码值，用于与存储中的正确值进行比较
//   clear: 指示在验证后是否清除验证码的布尔值，通常在验证成功后清除
// 返回值:
//   验证结果的布尔值，如果验证码正确则返回 true，否则返回 false
func (c *Captcha) Verify(idKey, value string, clear bool) bool {
	return c.store.Verify(idKey, value, clear)
}
