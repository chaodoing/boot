package method

// MethodPassword 是一个映射，用于存储密码加密方法及其对应的函数
// 它定义了不同的密码加密算法，如MD5、SHA1、SHA256和SHA512
// 这些方法可以接受一个字符串参数（密码），并根据需要接受额外的参数
// 每个方法都会返回一个字符串，通常是加密后的密码
var MethodPassword = map[string]func(string, ...string) string{
	"md5":    MD5Password,    // 使用MD5算法加密密码
	"sha1":   SHA1Password,   // 使用SHA1算法加密密码
	"sha256": SHA256Password, // 使用SHA256算法加密密码
	"sha512": SHA512Password, // 使用SHA512算法加密密码
}
