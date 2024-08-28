package database

import (
	`errors`
	`os`
	`strings`
	
	`gorm.io/gorm`
)

const PS = string(os.PathSeparator)

type Databases struct {
	names map[string]string
}

func (d Databases) Set(names map[string]string) {
	d.names = names
}

func (d Databases) Get(name string) (db *gorm.DB, err error) {
	if value, ok := d.names[name]; ok {
		var file = PS + strings.Trim(value, PS)
		return Connection(file)
	}
	return nil, errors.New("数据库名称对应的数据库配置文件不存在")
}
