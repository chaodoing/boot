package database

import (
	`github.com/chaodoing/boot/logger`
)

type (
	Config struct {
		Type     string        `json:"type" xml:"type" yaml:"type" ini:"TYPE" comment:"数据库类型"`
		Host     string        `json:"host" xml:"host" yaml:"host" ini:"HOST" comment:"数据库连接主机"`
		Port     uint16        `json:"port" xml:"port" yaml:"port" ini:"PORT" comment:"数据库连接端口"`
		Username string        `json:"username" xml:"username" yaml:"username" ini:"USERNAME" comment:"数据库连接用户"`
		Password string        `json:"password" xml:"password" yaml:"password" ini:"PASSWORD" comment:"数据库连接密码"`
		Database string        `json:"database" xml:"database" yaml:"database" ini:"NAME" comment:"数据库名称"`
		Charset  string        `json:"charset" xml:"charset" yaml:"charset" ini:"CHARSET" comment:"数据库连接字符集"`
		Logger   logger.Logger `json:"logger" xml:"logger" yaml:"logger" ini:"DATABASE_LOG" comment:"数据库日志数据 [1: silent 2: error 3: warn 4: info]"`
	}
)
