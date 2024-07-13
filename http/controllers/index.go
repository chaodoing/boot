package controllers

import (
	`github.com/chaodoing/boot/container`
	`github.com/chaodoing/boot/models`
	`github.com/chaodoing/boot/o`
	`github.com/chaodoing/boot/traits`
	`github.com/kataras/iris/v12`
)

type Index struct {
	*traits.Controller
}

func (i *Index) Get(ctx iris.Context, box container.Container) {
	db, err := box.Database()
	if err != nil {
		o.O(ctx, 3306, err.Error(), nil)
		return
	}
	var account models.Administrator
	err = db.Find(&account).Error
	if err != nil {
		o.O(ctx, 1, err.Error(), nil)
		return
	}
	o.O(ctx, 0, "OK", account)
	return
}

func (i *Index) GetAdmin(ctx iris.Context, box container.Container) {
	db, err := box.Database()
	if err != nil {
		o.O(ctx, 3306, err.Error(), nil)
		return
	}
	var admin models.Administrator
	err = db.Find(&admin).Error
	if err != nil {
		o.O(ctx, 1, err.Error(), nil)
		return
	}
	o.O(ctx, 0, "OK", admin)
	return
}
