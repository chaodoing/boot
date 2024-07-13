package container

import (
	`time`
	
	`github.com/chaodoing/boot/o`
	`github.com/iris-contrib/middleware/jwt`
	`github.com/kataras/iris/v12`
)

// Jwt 结构体用于管理 JWT（JSON Web Token）的中间件和相关配置。
type Jwt struct {
	jet        *jwt.Middleware // JWT 中间件实例
	secret     []byte          // 用于签名JWT的密钥
	duration   time.Duration   // JWT的过期时间
	Middleware iris.Handler    // JWT中间件的处理函数
}

// NewJwt 创建并返回一个新的 Jwt 实例，配置了指定的密钥和过期时间。
// @secret 密钥，用于JWT的签名。
// @duration JWT的过期时间。
// 返回 *Jwt，指向新创建的Jwt实例。
func NewJwt(secret string, duration time.Duration) *Jwt {
	jet := jwt.New(jwt.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		},
		SigningMethod: jwt.SigningMethodHS512,
		Expiration:    true,
		ErrorHandler: func(ctx iris.Context, err error) {
			o.O(ctx, 401, err.Error())
		},
	})
	return &Jwt{
		secret:     []byte(secret),
		duration:   duration,
		jet:        jet,
		Middleware: jet.Serve,
	}
}

// Get 从上下文中提取JWT，并返回其包含的"dial"字段的值。
// 如果JWT验证失败，将调用错误处理函数并返回nil。
// @ctx iris.Context，请求的上下文。
// 返回 interface{}，JWT中"dial"字段的值，如果验证失败则为nil。
func (e *Jwt) Get(ctx iris.Context) interface{} {
	if err := e.jet.CheckJWT(ctx); err != nil {
		e.jet.Config.ErrorHandler(ctx, err)
		return nil
	}
	token := ctx.Values().Get("jwt").(*jwt.Token)
	value := token.Claims.(jwt.MapClaims)
	return value["dial"]
}

// Tokenization 根据给定的数据创建一个新的JWT，并签名。
// @data 要包含在JWT中的数据，通常是一个映射。
// 返回签名后的JWT字符串，以及可能的错误。
func (e *Jwt) Tokenization(data interface{}) (string, error) {
	now := time.Now()
	token := jwt.NewTokenWithClaims(e.jet.Config.SigningMethod, jwt.MapClaims{
		"dial": data,
		"iat":  now.Unix(),
		"exp":  now.Add(e.duration).Unix(),
	})
	return token.SignedString(e.secret)
}
