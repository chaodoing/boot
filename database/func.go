package database

import (
	`fmt`
	`strings`
	
	`gorm.io/driver/mysql`
	`gorm.io/driver/postgres`
	`gorm.io/driver/sqlite`
	`gorm.io/driver/sqlserver`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
)

var drivers = map[string]func(c *Config, value *gorm.Config) (db *gorm.DB, err error){
	"mysql": func(c *Config, value *gorm.Config) (db *gorm.DB, err error) {
		schema := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
		return gorm.Open(mysql.Open(schema), value)
	},
	"sqlite": func(c *Config, value *gorm.Config) (db *gorm.DB, err error) {
		return gorm.Open(sqlite.Open(c.Database), value)
	},
	"postgres": func(c *Config, value *gorm.Config) (db *gorm.DB, err error) {
		schema := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable TimeZone=Asia/Shanghai", c.Host, c.Port, c.Username, c.Database, c.Password)
		return gorm.Open(postgres.Open(schema), value)
	},
	"sqlserver": func(c *Config, value *gorm.Config) (db *gorm.DB, err error) {
		schema := fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v", c.Username, c.Password, c.Host, c.Port, c.Database)
		return gorm.Open(sqlserver.Open(schema), value)
	},
}

func (c *Config) Connection() (db *gorm.DB, err error) {
	Log, err := c.Logger.Log()
	if err != nil {
		return
	}
	config := &gorm.Config{
		Logger: logger.New(Log, logger.Config{
			Colorful: false,
			LogLevel: c.Logger.GormLevel(),
		}),
		SkipDefaultTransaction: true,  // SkipDefaultTransaction 跳过默认事务
		FullSaveAssociations:   true,  // FullSaveAssociations 在创建或更新时，是否更新关联数据
		PrepareStmt:            true,  // PrepareStmt 是否禁止创建 prepared statement 并将其缓存
		AllowGlobalUpdate:      false, // AllowGlobalUpdate 是否允许全局 update/delete
		QueryFields:            true,  // QueryFields 执行查询时，是否带上所有字段
	}
	if fn, ok := drivers[strings.ToLower(c.Type)]; ok {
		return fn(c, config)
	} else {
		err = fmt.Errorf("不支持的数据库类型: %v", c.Type)
	}
	return
}
