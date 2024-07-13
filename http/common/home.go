package common

import (
	`github.com/chaodoing/boot/container`
	`github.com/chaodoing/boot/models`
	`github.com/chaodoing/boot/o`
	`github.com/kataras/iris/v12`
)

func Home(ctx iris.Context, box container.Container) {
	db, err := box.Database()
	if err != nil {
		o.O(ctx, 3306, err.Error())
		return
	}
	var account []models.Administrator
	if err := db.Find(&account).Error; err != nil {
		o.O(ctx, 3306, err.Error())
		return
	}
	o.O(ctx, account)
}
