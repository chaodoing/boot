package container

import (
	`fmt`
	`net/http`
	
	`github.com/kataras/iris/v12`
)

var Cors = func(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Headers", "Refresh-Token, Accept-Version, Authorization, Accept-Token, Language, Access-Control-Allow-Methods, Access-Control-Allow-Origin, Cache-Control, Content-Type, if-match, if-modified-since, if-none-match, if-unmodified-since, X-Requested-With")
	ctx.Header("Access-Control-Allow-Origin", fmt.Sprintf("%s", ctx.GetHeader("Origin")))
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	ctx.Header("Access-Control-Max-Age", "3600")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Expose-Headers", "Authorization, Accept-Token, Refresh-Token, Refresh-Expires")
	ctx.Header("Version", "1.0.0")
	ctx.Header("Author", "Neil")
	ctx.Header("Email", "chaodoing@live.com")
	if ctx.Method() == http.MethodOptions {
		ctx.StatusCode(204)
		return
	} else {
		ctx.Next()
	}
}
