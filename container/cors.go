package container

import (
	`time`
	
	`github.com/iris-contrib/middleware/cors`
)

var Cors = cors.New(cors.Options{
	AllowCredentials: true,
	AllowedOrigins:   []string{"*"},
	AllowedHeaders: []string{
		"Refresh-Token",
		"Accept-Version",
		"Authorization",
		"Accept-Token",
		"Language",
		"Access-Control-Allow-Methods",
		"Access-Control-Allow-Origin",
		"Cache-Control",
		"Content-Type",
		"if-match",
		"if-modified-since",
		"if-none-match",
		"if-unmodified-since",
		"X-Requested-With",
	},
	AllowedMethods: []string{
		"GET",
		"POST",
		"PUT",
		"DELETE",
		"PATCH",
		"OPTIONS",
	},
	ExposedHeaders: []string{
		"Authorization",
		"Accept-Token",
		"Refresh-Token",
		"Refresh-Expires",
	},
	MaxAge: int((24 * time.Hour).Seconds()), // 24 小时
})
