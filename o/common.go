package o

import (
	`github.com/chaodoing/boot/assets/vscode`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/mvc`
)

func Handle(app *iris.Application) {
	app.HandleDir("/vscode/static", vscode.Static)
	app.HandleDir("/vscode/monaco-editor", vscode.Monaco)
	app.OnAnyErrorCode(func(ctx iris.Context) {
		O(ctx, ctx.GetStatusCode(), iris.StatusText(ctx.GetStatusCode()))
	})
}

func Json[T Message[any] | Pagination[any]](data T) mvc.Result {
	return mvc.Response{
		ContentType: "application/json",
		Object:      data,
	}
}

func Xml[T Message[any] | Pagination[any]](data T) mvc.Result {
	return mvc.Response{
		ContentType: "application/xml",
		Object:      data,
	}
}

func Vscode[T Message[any] | Pagination[any]](data T, values ...any) mvc.Result {
	r := &Respond{}
	value, err := r.html(data)
	if err != nil {
		return mvc.Response{
			Err: err,
		}
	}
	result := mvc.Response{ContentType: "text/html", Content: []byte(value)}
	if len(values) > 0 {
		for _, value := range values {
			switch data := value.(type) {
			case string:
				result.Text = data
			case int:
				result.Code = data
			case error:
				result.Err = data
			}
		}
	}
	return result
}

func View(layout string, values ...any) mvc.Result {
	view := mvc.View{Layout: layout}
	for _, value := range values {
		switch data := value.(type) {
		case string:
			view.Name = data
		case int:
			view.Code = data
		case error:
			view.Err = data
		default:
			view.Data = data
		}
	}
	return view
}

func O(ctx iris.Context, data ...interface{}) {
	var code int
	if ctx.GetStatusCode() == 200 {
		code = 0
	}
	var result = Message[any]{
		Code:    code,
		Message: iris.StatusText(ctx.GetStatusCode()),
		Data:    nil,
	}
	for _, value := range data {
		switch val := value.(type) {
		case int:
			result.Code = val
		case string:
			result.Message = val
		default:
			result.Data = val
		}
	}
	_ = (&Respond{}).Negotiation(ctx, result)
}
