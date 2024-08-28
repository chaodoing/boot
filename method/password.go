package method

import (
	`crypto/md5`
	`crypto/sha1`
	`crypto/sha256`
	`crypto/sha512`
	`encoding/base64`
	`fmt`
)

// PasswordMD5Encrypt 使用MD5算法对密码进行加密。
// 它可以接受一个可选的盐值来增加安全性。
// 参数:
//   value: 需要加密的密码字符串。
//   salt: 可选的盐值，用于增加密码加密的安全性。
// 返回值:
//   pv: 加密后的密码字符串。
func PasswordMD5Encrypt(value string, salt ...string) (pv string) {
	// 检查是否有提供盐值，如果有，则在value前后加上盐值以增加安全性。
	if len(salt) > 0 {
		value = salt[0] + ":" + value + ":" + salt[0]
	}
	
	// 计算value的MD5哈希值。
	h := md5.Sum([]byte(value))
	
	// 将哈希结果转换为16进制字符串。
	pv = fmt.Sprintf("%x", h)
	
	// 返回加密后的密码。
	return
}

// PasswordSHA1Encrypt 用于对密码进行SHA1加密，并可添加盐值以增强安全性。
// 该函数接受一个字符串类型的密码和一个可变参数盐值。
// 如果提供了盐值，它将以指定的格式与密码一起被加密。
// 最终返回加密后的Base64编码字符串。
func PasswordSHA1Encrypt(s string, salt ...string) (p string) {
	// 检查是否有提供盐值，如果有，则按照特定格式将盐值和密码拼接起来。
	// 这里的盐值用于增加密码破解的难度，提高安全性。
	if len(salt) > 0 {
		s = salt[0] + ":" + s + ":" + salt[0]
	}
	
	// 创建SHA1哈希对象，用于执行SHA1加密算法。
	o := sha1.New()
	
	// 向哈希对象中写入字符串的字节序列。
	// 这一步是将密码（可能包含盐值）转换为字节序列进行哈希计算。
	o.Write([]byte(s))
	
	// 将哈希对象的最终结果编码为Base64字符串。
	// 使用Base64编码是为了将二进制数据安全地表示为ASCII字符，方便存储和传输。
	p = base64.StdEncoding.EncodeToString(o.Sum(nil))
	
	// 返回加密后的字符串。
	return
}

// PasswordSHA256Encrypt 使用SHA256算法对密码进行加密。
// 它接受一个字符串s作为待加密的密码，以及一个可选的盐值数组salt。
// 当提供盐值时，它会将盐值添加到密码的前后，以增加安全性。
// 函数返回加密后的Base64编码字符串p。
func PasswordSHA256Encrypt(s string, salt ...string) (p string) {
	// 检查是否有提供盐值
	if len(salt) > 0 {
		// 如果有盐值，则在密码前后添加盐值，并用冒号分隔，以增强加密的安全性
		s = salt[0] + ":" + s + ":" + salt[0]
	}
	// 创建SHA256哈希对象
	o := sha256.New()
	// 向哈希对象中写入字符串的字节序列
	o.Write([]byte(s))
	// 将哈希对象的最终结果编码为Base64字符串
	p = base64.StdEncoding.EncodeToString(o.Sum(nil))
	return
}

// PasswordSHA512Encrypt 使用 SHA512 算法和可选的盐值对密码进行加密。
// 参数 s 是待加密的密码字符串。
// 参数 salt 是一个字符串切片，用于接收可变长度的盐值参数，盐值用于增加密码哈希的安全性。
// 如果提供了盐值，它会在密码字符串前后各添加一次，并用冒号分隔，以增强哈希值的独特性。
// 返回值 p 是加密后的密码，以 base64 编码的字符串形式返回。
func PasswordSHA512Encrypt(s string, salt ...string) (p string) {
	// 检查是否有提供盐值，如果提供了，则按照规定格式调整字符串格式。
	if len(salt) > 0 {
		s = salt[0] + ":" + s + ":" + salt[0]
	}
	// 创建一个新的 SHA512 散列对象
	o := sha512.New()
	// 向散列对象中写入字符串 s 的字节表示
	o.Write([]byte(s))
	// 将散列对象的最终结果编码为 base64 字符串
	p = base64.StdEncoding.EncodeToString(o.Sum(nil))
	return
}

// PasswordMethods 是一个映射，用于存储密码加密方法及其对应的函数
// 它定义了不同的密码加密算法，如MD5、SHA1、SHA256和SHA512
// 这些方法可以接受一个字符串参数（密码），并根据需要接受额外的参数
// 每个方法都会返回一个字符串，通常是加密后的密码
var PasswordMethods = map[string]func(string, ...string) string{
	"md5":    PasswordMD5Encrypt,    // 使用MD5算法加密密码
	"sha1":   PasswordSHA1Encrypt,   // 使用SHA1算法加密密码
	"sha256": PasswordSHA256Encrypt, // 使用SHA256算法加密密码
	"sha512": PasswordSHA512Encrypt, // 使用SHA512算法加密密码
}
