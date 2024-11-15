package config

import (
	`encoding/xml`
	`os`
	
	`github.com/chaodoing/boot/auth/captcha`
	`github.com/chaodoing/boot/cache`
	`github.com/chaodoing/boot/database`
	`github.com/chaodoing/boot/logger`
	`github.com/gookit/goutil/fsutil`
	`gopkg.in/ini.v1`
)

type (
	Service struct {
		Favicon string        `json:"favicon" xml:"favicon" yaml:"Favicon" ini:"FAVICON" comment:"网站图标配置"`                                                                // Favicon 网站图标配置
		Host    string        `json:"host" xml:"host" yaml:"Host" ini:"HOST" comment:"监听主机"`                                                                                // Host 监听主机
		Port    uint16        `json:"port" xml:"port" yaml:"Port" ini:"PORT" comment:"监听端口"`                                                                                // Port 监听端口
		Cross   bool          `json:"cross" xml:"cross" yaml:"cross" ini:"CROSS" comment:"允许跨域访问"`                                                                        // Cross 允许跨域访问
		Logger  logger.Logger `json:"logger" xml:"logger" yaml:"logger" ini:"SERVICE_LOG" comment:"服务日志配置 level:[0->disable 1->fatal 2->error 3->warn 4->info 5->debug]"` // Logger 服务日志配置 服务日志配置 level:[0->disable 1->fatal 2->error 3->warn 4->info 5->debug]
	}
	Uploader struct {
		Path    string `json:"path" xml:"path" yaml:"path" ini:"PATH" comment:"访问地址"`
		Local   string `json:"local" xml:"local" yaml:"local" ini:"LOCAL" comment:"本地地址"`
		MaxSize int64  `json:"max_size" xml:"max_size" yaml:"max_size" ini:"MAX_SIZE" comment:"上传文件最大尺寸MB"`
	}
	Jwt struct {
		Secret string `json:"secret" xml:"secret" yaml:"secret" ini:"SECRET" comment:"JWT密钥"`
		Expire int64  `json:"expire" xml:"expire" yaml:"expire" ini:"EXPIRE" comment:"JWT过期时间"`
	}
)

type (
	View struct {
		Path         string `json:"path" xml:"path" yaml:"path" ini:"PATH" comment:"模板路径"`                                       // 模板路径
		LeftDelimit  string `json:"leftDelimit" xml:"delimit>left" yaml:"leftDelimit" ini:"LEFT_DELIMIT" comment:"模板左定界符"`     // 模板左定界符
		RightDelimit string `json:"rightDelimit" xml:"delimit>right" yaml:"rightDelimit" ini:"RIGHT_DELIMIT" comment:"模板右定界符"` // 模板右定界符
		Layout       string `json:"layout" xml:"layout" yaml:"layout" ini:"LAYOUT" comment:"模板布局"`                               // 默认模板布局
		Extension    string `json:"extension" xml:"extension" yaml:"extension" ini:"EXTENSION" comment:"模板后缀"`                   // 模板后缀
	}
	Config struct {
		XMLName  xml.Name        `xml:"root" json:"-" yaml:"-" ini:"-"`
		Service  Service         `json:"service" xml:"service" yaml:"service" ini:"SERVICE" comment:"网站服务配置"`
		Uploader Uploader        `json:"uploader" xml:"uploader" yaml:"uploader" ini:"UPLOADER" comment:"资源上传配置"`
		Database database.Config `json:"database" xml:"database" yaml:"database" ini:"DATABASE" comment:"数据库配置"`
		Captcha  captcha.Options `json:"captcha" xml:"captcha" yaml:"captcha" ini:"CAPTCHA" comment:"验证码配置"`
		Cache    cache.Config    `json:"cache" xml:"cache" yaml:"cache" ini:"CACHE" comment:"缓存配置"`
		Jwt      Jwt             `json:"jwt" xml:"jwt" yaml:"jwt" ini:"JWT" comment:"JWT配置"`
	}
)

func (c Config) LoadEnv() Config {
	// 构造开发环境配置文件的路径
	environment := os.ExpandEnv("${WORKDIR}/.env.${ENV}")
	// 构造通用配置文件的路径
	env := os.ExpandEnv("${WORKDIR}/.env")
	if fsutil.FileExist(environment) {
		f, err := ini.Load(environment)
		if err == nil {
			_ = f.MapTo(&c)
		}
	} else if fsutil.FileExist(env) {
		f, err := ini.Load(env)
		if err == nil {
			_ = f.MapTo(&c)
		}
	}
	return c
}
