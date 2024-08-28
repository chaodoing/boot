package traits

import (
	`github.com/chaodoing/boot/method`
	`github.com/gookit/validate`
)

// LoginValidate 结构体代表登录验证功能，其中的 PasswordValidate 方法用于验证密码的有效性。
type LoginValidate struct {
	Username string `json:"username" xml:"username" validate:"required"`
	Password string `json:"password" xml:"password" validate:"passwordValidate"`
}

// PasswordValidate 方法用于验证给定的密码是否有效。
// 该方法调用了一个独立的验证方法（method.PasswordValidate）来执行实际的验证逻辑。
// 参数:
//   value string: 待验证的密码字符串。
// 返回值:
//   bool: 如果密码有效，则返回true；否则返回false。
func (l LoginValidate) PasswordValidate(value string) bool {
	return method.PasswordValidate(value)
}

// ConfigValidation 配置验证规则
// 该方法为 LoginValidate 类的成员函数，专门用于设置登录验证场景下的特定验证规则
// 参数 v: 一个指向 validate.Validation 类型的指针，用于执行验证逻辑
func (l LoginValidate) ConfigValidation(v *validate.Validation) {
	// 设置 "login" 场景下的验证规则，指定需要验证的字段为 "Username" 和 "Password"
	// 这一步确保了在 "login" 场景中，用户名和密码是必须通过验证的字段
	v.WithScenes(validate.SValues{
		"login": []string{"Username", "Password"},
	})
}

// Messages 返回登录验证相关的错误信息映射。
// 该函数定义了验证规则中各字段在不满足条件时的错误提示信息。
// 返回值是一个映射，其中键是需要验证的字段名，值是对应的错误信息。
func (l LoginValidate) Messages() map[string]string {
	return validate.MS{
		"required":         "{field}不能为空",                                                                                                                  // 提示信息用于required验证规则，表示字段不能为空。
		"passwordValidate": "{field}长度不小于8个字符 包含至少一个数字 包含至少一个小写字母 包含至少一个大写字母 包含至少一个特殊字符 !@#~$%^&*()./-=\"';,+|_", // 提示信息用于passwordValidate规则，详细说明了密码的复杂度要求。
	}
}

// Translates 返回字段名称到对应的中文显示名称的映射。
// 该函数用于为登录验证中的每个字段提供友好的显示名称，促进国际化和本地化支持。
// 该函数没有输入参数，返回值是map[string]字符串类型，包含字段名到显示名的映射
func (l LoginValidate) Translates() map[string]string {
	return validate.MS{
		"Username": "登录账号",
		"Password": "登录密码",
	}
}
