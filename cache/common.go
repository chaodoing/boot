package cache

import (
	`strings`
	
	`github.com/go-redis/redis`
)

// New 创建并返回一个新的Cache实例。
// 参数config是一个Config指针，用于配置Cache的行为和设置。
// 返回值是一个指向Cache实例的指针。
func New(rdx *redis.Client, prefixes ...string) (*Cache, error) {
	var prefix = "cache"
	if len(prefixes) > 0 {
		prefix = prefixes[0]
	}
	if err := rdx.Ping().Err(); err != nil {
		return nil, err
	}
	// 根据提供的配置信息初始化并返回一个新的Cache实例。
	return &Cache{
		rdx:    rdx,
		prefix: prefix,
	}, nil
}

// NewGroup 根据给定的配置信息和前缀创建一个新的Group实例。
// config: 配置信息，用于连接Redis。
// prefixes: 可变参数，用于设置键的前缀。如果提供单个前缀，则直接使用；如果提供两个前缀，则使用它们来创建一个冒号分隔的字符串作为前缀。
// 返回一个新的Group实例，以及可能的错误。
func NewGroup(rdx *redis.Client, prefixes ...string) (*Group, error) {
	// 默认前缀为"group"
	var prefix = "cache:group"
	// 如果提供了单个前缀，则直接使用它
	if len(prefixes) == 1 {
		prefix = prefixes[0]
	} else if len(prefixes) >= 2 {
		// 如果提供了两个前缀，则将它们合并为一个冒号分隔的字符串
		prefix = strings.Join(prefixes, ":")
	}
	// 检查Redis连接是否正常
	if err := rdx.Ping().Err(); err != nil {
		return nil, err
	}
	// 创建并返回一个新的Group实例，其中包含一个使用给定配置和前缀初始化的Cache实例
	// 根据提供的配置信息初始化并返回一个新的Cache实例。
	return &Group{
		&Cache{
			rdx:    rdx,
			prefix: prefix,
		},
	}, nil
}
