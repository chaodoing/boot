package database

import (
	`encoding/json`
	`encoding/xml`
	`errors`
	`fmt`
	`os`
	`strings`
	
	`github.com/gookit/goutil/fsutil`
	`gopkg.in/yaml.v2`
	`gorm.io/gorm`
)

// xmlConnect 根据配置的 XML 数据建立数据库连接。
//
// 参数:
//   config - Config 类型，用于存储解析后的 XML 配置信息。
//   data - byte 类型切片，代表待解析的 XML 配置数据。
//
// 返回值:
//   *gorm.DB - 数据库连接实例。
//   error - 解析配置或建立数据库连接时可能出现的错误。
//
// 该函数首先尝试将 XML 数据解析到 config 对象中，如果解析失败，则直接返回错误。
// 解析成功后，调用 config 对象的 Connection 方法建立并返回数据库连接。
func xmlConnect(config Config, data []byte) (db *gorm.DB, err error) {
	// 尝试解析 XML 数据到 config 对象
	err = xml.Unmarshal(data, &config)
	// 如果解析失败，返回错误
	if err != nil {
		return
	}
	// 解析成功，返回建立的数据库连接
	return config.Connection()
}

// jsonConnect 根据配置信息初始化数据库连接。
// 它首先尝试解析传入的JSON数据为Config结构体，然后使用该配置信息建立数据库连接。
// 参数:
//   config - Config类型的变量，用于存储解析后的配置信息。
//   data - 包含配置信息的JSON格式字节切片。
// 返回值:
//   *gorm.DB - 成功连接数据库后返回的数据库连接对象。
//   error - 解析JSON数据或建立数据库连接时可能出现的错误。
func jsonConnect(config Config, data []byte) (db *gorm.DB, err error) {
	// 尝试解析JSON数据到Config结构体
	err = json.Unmarshal(data, &config)
	// 如果解析出错，则返回错误信息
	if err != nil {
		return
	}
	// 返回使用Config信息建立的数据库连接
	return config.Connection()
}

// ymlConnect 根据配置文件初始化数据库连接。
// 它首先尝试解析传入的YAML配置数据，然后根据解析后的配置建立数据库连接。
// 参数:
//   config - Config 类型的对象，用于存储解析后的配置信息。
//   data - 二进制形式的YAML配置数据。
// 返回值:
//   *gorm.DB - 数据库连接对象。
//   error - 解析配置或建立数据库连接时可能出现的错误。
func ymlConnect(config Config, data []byte) (db *gorm.DB, err error) {
	// 使用yaml.Unmarshal解析配置数据到config对象
	err = yaml.Unmarshal(data, &config)
	// 如果解析过程中出现错误，则返回错误信息
	if err != nil {
		return
	}
	// 返回根据配置信息建立的数据库连接
	return config.Connection()
}

// Connection 根据配置名称初始化数据库连接。
// 参数name可以是配置文件的名称或Config类型的实例。
// 返回值db是gorm.DB类型的数据库连接实例，err是可能发生的错误。
func Connection(name interface{}) (db *gorm.DB, err error) {
	var dir string
	if dir = os.Getenv("CONFIG_DIR"); strings.EqualFold(dir, "") {
		err = errors.New("environment variable CONFIG_DIR directory not found")
		return
	}
	// 初始化配置结构体
	config := Config{}
	// 根据name的类型进行处理
	switch a := name.(type) {
	case string:
		// 检查传入的字符串是否是文件路径，并根据文件扩展名进行不同类型的配置加载
		if fsutil.FileExist(os.ExpandEnv(a)) {
			as := strings.Split(a, ".")
			ext := as[len(as)-1]
			data := fsutil.GetContents(os.ExpandEnv(a))
			// 根据文件扩展名，调用不同的连接函数
			switch ext {
			case "xml":
				return xmlConnect(config, data)
			case "json":
				return jsonConnect(config, data)
			case "yaml", "yml":
				return ymlConnect(config, data)
			}
		}
		// 如果传入的字符串不是文件路径，尝试按照默认配置文件路径加载
		xmlConfig := fmt.Sprintf("%s/%s.xml", dir, a)
		if fsutil.FileExist(xmlConfig) {
			data := fsutil.GetContents(xmlConfig)
			return xmlConnect(config, data)
		}
		jsonConfig := fmt.Sprintf("%s/%s.json", dir, a)
		if fsutil.FileExist(xmlConfig) {
			data := fsutil.GetContents(jsonConfig)
			return jsonConnect(config, data)
		}
		yamlConfig := fmt.Sprintf("%s/%s.yaml", dir, a)
		if fsutil.FileExist(yamlConfig) {
			data := fsutil.GetContents(yamlConfig)
			return ymlConnect(config, data)
		}
		ymlConfig := fmt.Sprintf("%s/%s.yml", dir, a)
		if fsutil.FileExist(yamlConfig) {
			data := fsutil.GetContents(ymlConfig)
			return ymlConnect(config, data)
		}
	case Config:
		// 如果name是Config类型，则直接使用其Connection方法初始化数据库连接
		return a.Connection()
	}
	// 如果无法处理name的类型，则返回错误提示配置文件未找到
	err = errors.New("config file not found")
	return
}
