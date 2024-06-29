package task

import (
	`fmt`
	`reflect`
)

type Events struct {
	listeners map[string][]interface{}
}

func NewEvent() *Events {
	return &Events{listeners: make(map[string][]interface{})}
}

func (e *Events) AddEventListener(name string, fn interface{}) {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("not a function")
	}
	e.listeners[name] = append(e.listeners[name], fn)
}

func (e *Events) callFunctionWithReflection(fn interface{}, params ...interface{}) {
	// 获取函数的类型信息
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("not a function")
	}
	
	// 获取函数的值信息
	fnValue := reflect.ValueOf(fn)
	
	// 将参数的interface{}切片转换为reflect.Value切片
	args := make([]reflect.Value, len(params))
	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}
	
	// 调用函数
	fnValue.Call(args)
}

func (e *Events) Trigger(name string, args ...interface{}) {
	if values, ok := e.listeners[name]; ok {
		for _, fn := range values {
			fmt.Println(fn)
			e.callFunctionWithReflection(fn, args...)
		}
	} else {
		panic("事件不存在")
	}
}

func (e *Events) Keep() {

}

func (e *Events) Wake() {

}

func (e *Events) Clear() {

}
