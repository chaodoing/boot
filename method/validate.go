package method

import (
	`regexp`
)

// ValidatePassword 该函数用于验证密码是否符合以下要求：
//   长度不小于8个字符。
//   包含至少一个数字。
//   包含至少一个小写字母。
//   包含至少一个大写字母。
//   包含至少一个特殊字符（如 [!@#~$%^&*()./-="';,+|_] ）。
//   函数使用正则表达式对传入的密码进行匹配，如果密码不满足其中任何一项要求，则返回false，否则返回true。
//   This function validates that the password is at least 8 characters long and contains at least one digit, one lowercase letter, one uppercase letter, and one special character.
//   Parameters:
//      value - 要验证的密码值
//   Return value:
//      pass - 显示密码是否符合要求。满足要求为True，不满足要求为false。
func ValidatePassword(value string) (pass bool) {
	// 初始化密码验证结果为true
	pass = true
	// 检查密码长度，如果小于8，则不满足要求。
	if len(value) < 8 {
		pass = false
	}
	// 定义正则表达式模式
	var (
		num    = `[0-9]{1}`                    // Pattern for matching at least one digit.
		a_z    = `[a-z]{1}`                    // Pattern for matching at least one lowercase letter.
		A_Z    = `[A-Z]{1}`                    // Pattern for matching at least one uppercase letter.
		symbol = `[!@#~$%^&*()./-="';,+|_]{1}` // Pattern for matching at least one special character.
	)
	
	// 检查是否包含至少一个数字，如果不包含，则不满足要求。
	if b, err := regexp.MatchString(num, value); !b || err != nil {
		pass = false
	}
	
	// 检查是否包含至少一个小写字母，如果不包含，则不满足要求。
	if b, err := regexp.MatchString(a_z, value); !b || err != nil {
		pass = false
	}
	
	// 检查是否包含至少一个大写字母，如果不包含，则不满足要求。
	if b, err := regexp.MatchString(A_Z, value); !b || err != nil {
		pass = false
	}
	
	// 检查是否包含至少一个特殊字符，如果不包含，则不满足要求。
	if b, err := regexp.MatchString(symbol, value); !b || err != nil {
		pass = false
	}
	
	// 返回密码验证结果
	return
}
