package config

import (
	`encoding/json`
	`encoding/xml`
	`os`
	`strings`
	
	`github.com/chaodoing/boot/cache`
	`github.com/chaodoing/boot/database`
	`github.com/chaodoing/boot/logger`
	`github.com/gookit/goutil/fsutil`
	`github.com/kataras/iris/v12/x/errors`
	`gopkg.in/ini.v1`
	`gopkg.in/yaml.v2`
)

func TestENV() error {
	var (
		ENV        = os.Getenv("ENV")
		WORKDIR    = os.Getenv("WORKDIR")
		CONFIG_DIR = os.Getenv("CONFIG_DIR")
		LOG_DIR    = os.Getenv("LOG_DIR")
	)
	if strings.EqualFold("", CONFIG_DIR) {
		return errors.New("environment variable CONFIG_DIR not found")
	}
	if strings.EqualFold("", LOG_DIR) {
		return errors.New("environment variable LOG_DIR not found")
	}
	if strings.EqualFold("", WORKDIR) {
		return errors.New("environment variable WORKDIR not found")
	}
	if strings.EqualFold("", ENV) {
		return errors.New("environment variable ENV not found")
	}
	return nil
}

// INIWriter 将给定的数据结构转换为INI格式的配置文件。
// 它首先创建一个空的INI配置对象，然后通过反射从数据结构中填充配置项。
// 最后，它将配置保存到两个文件中，一个是通用的环境配置文件，另一个是特定环境的配置文件。
// 参数:
//   data - 待转换为INI格式的配置数据结构的接口。
// 返回值:
//   error - 如果在处理过程中发生错误，则返回该错误。
func INIWriter(data Config) (err error) {
	// 创建一个空的INI配置对象
	env := ini.Empty()
	
	// 通过反射从数据结构中填充INI配置对象
	err = ini.ReflectFrom(env, &data)
	if err != nil {
		return err
	}
	fileName := os.ExpandEnv("${WORKDIR}/.env")
	fileEnv, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fileEnv.Close()
	// 将配置保存到通用环境配置文件中
	// 使用os.ExpandEnv替换环境变量，确保文件路径的正确性
	err = env.SaveTo(fileName)
	if err != nil {
		return err
	}
	fileName = os.ExpandEnv("${WORKDIR}/.env.${ENV}")
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	// 将配置保存到特定环境配置文件中
	// 这允许根据不同的环境（如开发、测试、生产）使用不同的配置
	err = env.SaveTo(fileName)
	if err != nil {
		return err
	}
	// 如果没有错误发生，返回nil表示操作成功
	return nil
}

// Default 函数返回一个默认的配置结构体实例，该实例包含了服务、上传、数据库、缓存和JWT的配置信息。
// 这些配置用于初始化和配置应用程序的不同组件，为快速启动项目提供了一种预设的配置方式。
func Default() Config {
	return Config{
		Service: Service{
			Favicon: "favicon.ico", // 默认的网站图标路径
			Host:    "0.0.0.0",     // 服务监听的主机地址
			Port:    9898,          // 服务监听的端口号
			Cross:   true,          // 默认开启跨域访问控制
			Logger: logger.Logger{
				Stdout: true,                     // 日志输出到标准输出
				Level:  5,                        // 日志级别为5，详细日志
				File:   "${LOG_DIR}/iris-%F.log", // 日志文件路径，包含日期格式
			},
		},
		Uploader: Uploader{
			Path:    "/upload",                     // 上传接口的路径
			Local:   "${WORKDIR}/resources/upload", // 上传文件在本地存储的路径
			MaxSize: 20,                            // 最大上传文件大小，单位为MB
		},
		Database: database.Config{
			Type:     "mysql",     // 数据库类型为MySQL
			Host:     "127.0.0.1", // 数据库服务器地址
			Port:     3306,        // 数据库服务器端口
			Username: "root",      // 数据库用户名
			Password: "123.com",   // 数据库密码
			Database: "test",      // 数据库名称
			Charset:  "utf8mb4",   // 数据库字符集
			Logger: logger.Logger{
				Stdout: true,                      // 数据库操作的日志输出到标准输出
				Level:  4,                         // 数据库操作的日志级别为4，一般日志
				File:   "${LOG_DIR}/mysql-%F.log", // 数据库操作的日志文件路径，包含日期格式
			},
		},
		Cache: cache.Config{
			Host:     "127.0.0.1",   // 缓存服务器地址
			Port:     6379,          // 缓存服务器端口
			Password: "123.com",     // 缓存服务器密码
			Index:    0,             // 缓存服务器的默认数据库索引
			TTL:      7 * 24 * 3600, // 缓存项的默认过期时间，单位为秒
		},
		Jwt: Jwt{
			Secret: "192.168.cc",  // JWT的签名密钥
			Expire: 7 * 24 * 3600, // JWT的默认过期时间，单位为秒
		},
	}
}

// Json 从指定的JSON文件中加载配置信息。
// 参数file是JSON配置文件的路径。
// 返回值data是解析后的配置信息。
// 如果解析过程中发生错误，错误返回值err将非空。
func Json(file string) (data Config, err error) {
	// 读取JSON文件的内容
	value := fsutil.GetContents(file)
	// 解析JSON内容到配置结构体
	err = json.Unmarshal(value, &data)
	return
}

// Xml 从指定的XML文件中加载配置信息。
// 参数file是XML配置文件的路径。
// 返回值data是解析后的配置信息。
// 如果解析过程中发生错误，错误返回值err将非空。
func Xml(file string) (data Config, err error) {
	// 读取XML文件的内容
	value := fsutil.GetContents(file)
	// 解析XML内容到配置结构体
	err = xml.Unmarshal(value, &data)
	return
}

// Yaml 从指定的YAML文件中加载配置信息。
// 参数file是YAML配置文件的路径。
// 返回值data是解析后的配置信息。
// 如果解析过程中发生错误，错误返回值err将非空。
func Yaml(file string) (data Config, err error) {
	// 读取YAML文件的内容
	value := fsutil.GetContents(file)
	// 解析YAML内容到配置结构体
	err = yaml.Unmarshal(value, &data)
	return
}
