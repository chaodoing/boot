package o

import (
	`github.com/gookit/validate`
)

// Validate 对给定的值进行验证。
//
// 参数:
//   value - 需要验证的值，可以是任何类型。
//   s     - 验证规则的参数，根据不同的验证规则，可以接收不同的参数。
//
// 返回值:
//   ok    - 表示验证是否通过的布尔值，通过返回true，否则返回false。
//   validation - 验证结果对象，包含验证过程中的详细信息。
//
// 说明:
//   该函数使用validate.Struct方法对传入的值进行结构化验证，首先检查value是否符合指定的验证规则。
//   如果value通过验证，则返回true和包含验证信息的validation对象。如果未通过验证，则返回false和包含验证详细信息的validation对象。
//   支持通过可变参数s指定不同的验证规则，提高了函数的灵活性和可扩展性。
func Validate(value interface{}, s ...string) (ok bool, validation *validate.Validation) {
	v := validate.Struct(value)
	if v.Validate(s...) {
		return true, v
	}
	return false, v
}

// MapValidate 对给定的数据映射进行验证。
//
// 参数:
// - value: 待验证的值，是一个字符串键任何类型值的映射。
// - rules: 验证规则映射，其中键是待验证的字段，值是该字段的验证规则字符串。
// - messages: 自定义错误消息映射，键是字段名，值是该字段在验证失败时显示的错误消息。
// - translates: 字段翻译映射，用于在错误消息中将字段名翻译成更易读的形式。
// - scenes: 指定的验证场景，用于在不同的业务场景中应用不同的验证规则。
// - s: 可变参数，用于指定在特定场景下需要验证的字段。
//
// 返回值:
// - ok: 如果验证成功，则返回true；否则返回false。
// - validation: 无论验证成功还是失败，都会返回一个包含验证结果的*validate.Validation对象。
func MapValidate(value map[string]any, rules map[string]string, messages map[string]string, translates map[string]string, scenes validate.SValues, s ...string) (ok bool, validation *validate.Validation) {
	// 创建一个验证对象，用于后续的验证操作。
	v := validate.Map(value)
	
	// 遍历规则映射，为每个字段添加验证规则。
	for field, args := range rules {
		v.StringRule(field, args)
	}
	
	// 添加自定义错误消息，提高验证失败时错误信息的可读性。
	v.AddMessages(messages)
	
	// 添加字段翻译，使错误消息中的字段名更易于理解。
	v.AddTranslates(translates)
	
	// 应用不同的验证场景，以便在不同的业务场景中使用不同的验证规则。
	v.WithScenes(scenes)
	
	// 执行验证，如果验证成功，则返回true和验证对象；否则返回false和验证对象。
	if v.Validate(s...) {
		return ok, v
	}
	
	return false, v
}
