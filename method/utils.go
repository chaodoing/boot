package method

import (
	`github.com/google/uuid`
)

// UUID 生成一个唯一的UUID字符串。
//
// 参数: 无
//
// 返回值: 一个字符串，代表一个唯一的UUID。
func UUID() string {
	return uuid.New().String()
}
