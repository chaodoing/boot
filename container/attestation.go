package container

import (
	`encoding/json`
	`errors`
	`strings`
	`time`
	
	`github.com/chaodoing/boot/method`
	`github.com/go-redis/redis`
	`github.com/kataras/iris/v12`
)

const (
	Basic  = "Basic "
	Bearer = "Bearer "
)

type (
	Attestation struct {
		rdx        *redis.Client
		timeToLive time.Duration
		prefix     string
	}
)

func trim(s string) string {
	return strings.TrimPrefix(strings.TrimPrefix(s, Basic), Bearer)
}

// expire 获取有效时间
func (a Attestation) expire() (date time.Time) {
	date = time.Now().Add(a.timeToLive)
	return
}

// token 生成用户认证字符串
func (a Attestation) token() (token string) {
	return method.UUID()
}

// current 当前认证字符串
func (a Attestation) current(ctx iris.Context) (token string) {
	if token = ctx.GetHeader("Access-Token"); !strings.EqualFold(token, "") {
		token = trim(token)
	}
	if token = ctx.GetHeader("Authorization"); !strings.EqualFold(token, "") {
		token = trim(token)
	}
	return
}

func (a Attestation) header(ctx iris.Context, token ...string) {
	ctx.Header("Access-Control-Allow-Headers", "Refresh-Token, Accept-Version, Authorization, Access-Token, Language, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Cache-Control, Content-Type, if-match, if-modified-since, if-none-match, if-unmodified-since, X-Requested-With")
	ctx.Header("Access-Control-Expose-Headers", "Authorization, Access-Token, Refresh-Token, Refresh-Expires")
	ctx.Header("Refresh-Expires", a.expire().Format("2006-01-02 15:04:05"))
	if len(token) > 0 {
		ctx.Header("Refresh-Token", token[0])
	}
}

func (a Attestation) notRefresh(ctx iris.Context) bool {
	var refresh = ctx.GetHeader("Refresh-Token")
	return strings.EqualFold(refresh, "0") || strings.EqualFold(refresh, "false") || strings.EqualFold(refresh, "off")
}

// GET 读取用户信息
func (a Attestation) GET(ctx iris.Context, data interface{}) (err error) {
	if token := a.current(ctx); !strings.EqualFold(token, "") {
		var content string
		content, err = a.rdx.Get(a.prefix + token).Result()
		if err != nil {
			return
		}
		err = json.Unmarshal([]byte(content), &data)
		if err != nil {
			return
		}
		if a.notRefresh(ctx) {
			_, err = a.rdx.Set(a.prefix+token, content, a.timeToLive).Result()
			if err != nil {
				return
			}
			a.header(ctx, "false")
		} else {
			_, err = a.rdx.Del(a.prefix + token).Result()
			if err != nil {
				return
			}
			token = a.token()
			_, err = a.rdx.Set(a.prefix+token, content, a.timeToLive).Result()
			if err != nil {
				return
			}
			a.header(ctx, token)
		}
	} else {
		err = errors.New("用户认证字符串是空的")
	}
	return
}

// SET 设置用户信息
func (a Attestation) SET(ctx iris.Context, data interface{}) (err error) {
	token := a.token()
	var bit []byte
	bit, err = json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = a.rdx.Set(a.prefix+token, string(bit), a.timeToLive).Result()
	if err != nil {
		return
	}
	a.header(ctx, token)
	return
}

// ALTER 修改用户信息
func (a Attestation) ALTER(ctx iris.Context, data interface{}) (err error) {
	if token := a.current(ctx); !strings.EqualFold(token, "") {
		if !a.notRefresh(ctx) {
			token = a.token()
		}
		var bit []byte
		bit, err = json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = a.rdx.Set(a.prefix+token, string(bit), a.timeToLive).Result()
		if err != nil {
			return
		}
		a.header(ctx, token)
		return
	} else {
		return errors.New("用户认证字符串是空的")
	}
}

// DEL 清理用户信息
func (a Attestation) DEL(ctx iris.Context) (err error) {
	if token := a.current(ctx); !strings.EqualFold(token, "") {
		_, err = a.rdx.Del(a.prefix + token).Result()
		if err != nil {
			return
		}
		return
	} else {
		return errors.New("用户认证字符串是空的")
	}
}
