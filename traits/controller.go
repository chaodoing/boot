package traits

import (
	`github.com/chaodoing/boot/o`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/mvc`
)

type Controller struct {
	Ctx iris.Context
}

func (c *Controller) BeforeActivation(m mvc.BeforeActivation) {}

func (c *Controller) HandleHTTPError(ctx iris.Context) {}

func (c *Controller) BeginRequest(ctx iris.Context) {}

func (c *Controller) EndRequest(ctx iris.Context) {}

func (c *Controller) Get(ctx iris.Context) mvc.Response {
	return mvc.Response{
		Code: iris.StatusOK,
		Text: "数据列表",
		Object: o.Message[any]{
			Code:    0,
			Message: "OK",
			Data:    nil,
		},
	}
}

func (c *Controller) Post(ctx iris.Context) mvc.Response {
	return mvc.Response{
		Code: iris.StatusOK,
		Text: "创建数据",
		Object: o.Message[any]{
			Code:    0,
			Message: "OK",
			Data:    nil,
		},
	}
}

func (c *Controller) Put(ctx iris.Context) mvc.Response {
	return mvc.Response{
		Code: iris.StatusOK,
		Text: "修改数据",
		Object: o.Message[any]{
			Code:    0,
			Message: "OK",
			Data:    nil,
		},
	}
}

func (c *Controller) Delete(ctx iris.Context) mvc.Response {
	return mvc.Response{
		Code: iris.StatusOK,
		Text: "删除数据",
		Object: o.Message[any]{
			Code:    0,
			Message: "OK",
			Data:    nil,
		},
	}
}
