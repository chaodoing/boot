package task

import (
	`errors`
	`reflect`
)

type Events struct {
	listeners map[string][]interface{}
}

// NewEvent 创建一个新的事件对象。
// 返回一个指向Events实例的指针，该实例初始化了一个空的监听器映射。
// 这允许用户注册对特定事件类型的监听器，以便在事件发生时进行处理。
func NewEvent() *Events {
	// 初始化一个空的监听器映射，键为事件类型，值为监听器列表。
	return &Events{listeners: make(map[string][]interface{})}
}

// AddEventListener 为指定的事件添加一个监听函数。
// name: 事件的名称。
// fn: 监听函数，必须是一个函数类型的值。
// 该方法会检查传入的fn是否为函数类型，如果不是，则抛出panic。
// 它将监听函数添加到对应事件的监听函数列表中。
func (e *Events) AddEventListener(name string, fn interface{}) {
	// 使用反射检查fn是否为函数类型
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("not a function")
	}
	// 将监听函数添加到事件列表中，如果该事件不存在，则先初始化一个空的函数切片
	e.listeners[name] = append(e.listeners[name], fn)
}

// Listeners 返回当前事件处理器中注册的所有监听器的名称列表。
//
// 该方法遍历内部监听器映射表，收集并返回所有监听器的名称。
// 这对于需要了解系统中当前注册了哪些监听器的场合非常有用，比如在调试或日志记录中。
func (e *Events) Listeners() []string {
	// 初始化一个空的字符串切片，用于存储监听器名称。
	var names = make([]string, 0)
	// 遍历监听器映射表，获取每个监听器的名称。
	for i, _ := range e.listeners {
		// 将监听器名称添加到名称列表中。
		names = append(names, i)
	}
	// 返回包含所有监听器名称的列表。
	return names
}

// callFunction 调用传入的函数，并传入参数列表。
// fn 是待调用的函数接口，params 是函数的参数列表，可变长。
// 函数返回执行结果的错误信息。
func (e *Events) callFunction(fn interface{}, params ...interface{}) error {
	// 获取fn的类型信息，用于检查是否为函数类型。
	fnType := reflect.TypeOf(fn)
	// 如果fn的类型不是函数类型，则返回错误。
	if fnType.Kind() != reflect.Func {
		return errors.New("not a function")
	}
	// 获取fn的值信息，用于后续的函数调用。
	fnValue := reflect.ValueOf(fn)
	// 根据params的长度创建一个反射值切片，用于存放函数调用的参数。
	args := make([]reflect.Value, len(params))
	// 遍历params，将每个参数转换为反射值，并存入args切片。
	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}
	// 使用反射调用fnValue对应的函数，传入args作为参数。
	fnValue.Call(args)
	// 函数调用成功，返回nil。
	return nil
}

// Trigger 触发一个事件及其相关的监听函数。
// 参数 name 代表事件的名称，args 是传递给监听函数的可变参数。
// 如果事件存在监听函数，则依次调用这些监听函数，并传递 args 参数。
// 如果事件不存在监听函数，则返回一个错误。
// 返回值是操作的结果，如果成功触发事件则为 nil，否则为一个错误对象。
func (e *Events) Trigger(name string, args ...interface{}) error {
	// 检查是否存在名为 name 的事件监听函数
	if values, ok := e.listeners[name]; ok {
		// 遍历所有监听函数并调用它们
		for _, fn := range values {
			if err := e.callFunction(fn, args...); err != nil {
				return err
			}
		}
		// 成功触发事件，返回 nil
		return nil
	} else {
		// 事件不存在，返回错误
		return errors.New("no event listener")
	}
}
