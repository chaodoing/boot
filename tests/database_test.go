package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/database`
	`github.com/chaodoing/boot/logger`
	`github.com/chaodoing/boot/models`
)

func TestDatabase(t *testing.T) {
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
		t.Error(err)
	}
	var account models.Account
	err = db.Table("account").First(&account).Error
	if err != nil {
		t.Error(err)
	}
	t.Log(account)
}
