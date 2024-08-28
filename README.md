# 框架工具包整合

## 数据库整合后支持独立调用

```go
package main
import (
	`github.com/chaodoing/boot/database`
	`github.com/chaodoing/boot/logger`
)
func main() {
	env := database.Config{
		Type:     "mysql",
		Host:     "192.168.33.10",
		Port:     3306,
		Username: "root",
		Password: "123.com",
		Database: "admin",
		Charset:  "utf8mb4",
		Logger: logger.Logger{
			Stdout: true,
			Level:  4,
			File:   "./logs/mysql-%F.log",
		},
	}
	db, err := env.Connection()
	if err != nil {
		panic(err)
	}
	var account map[string]interface{}
	err = db.Table("account").Where("id = ?", 1).First(&account).Error
	if err != nil {
		panic(err)
	}
}
```

```go
package main
import (
	`github.com/chaodoing/boot/cache`
)
func main() {
	config := cache.Config{
		Host:     "127.0.0.1",
		Port:     6379,
		Password: "123.com",
		Index:    0,
		TTL:      0,
	}
	g, err := cache.NewGroup(&config, "cache-group", "authorized")
	if err != nil {
		panic(err)
	}
}
```