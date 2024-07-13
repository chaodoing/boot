package cache

import (
	`errors`
	`fmt`
	`strings`
	`time`
)

// Group 是一个基于Cache的分组结构，它通过前缀区分不同的分组。
type Group struct {
	*Cache
}

// hasGroupName 检查当前分组是否有有效的分组名。
// 它通过判断前缀是否包含至少一个冒号来确定。
func (g *Group) hasGroupName() bool {
	return len(strings.Split(g.prefix, ":")) >= 2
}

// All 返回当前分组下所有键值对。
// 它通过匹配前缀来获取键值，并移除键的前缀以得到最终的结果。
func (g *Group) All() map[string]string {
	var result = map[string]string{}
	names := g.rdx.Keys(fmt.Sprintf("%s:*", g.prefix)).Val()
	key := fmt.Sprintf("%s:", g.prefix)
	for _, name := range names {
		result[strings.TrimPrefix(name, key)] = g.rdx.Get(name).Val()
	}
	return result
}

// Name 为当前分组设置一个新的分组名前缀。
// 这允许在同一个Cache实例中对不同的数据集进行分组操作。
func (g *Group) Name(name string) *Group {
	g.prefix = fmt.Sprintf("%s:%s", g.prefix, name)
	return g
}

// Set 在当前分组中设置键值对。
// 如果没有有效的分组名，则操作失败。
func (g *Group) Set(key string, value interface{}, ttl ...int) error {
	if g.hasGroupName() {
		return g.Cache.Set(key, value, ttl...)
	}
	return errors.New("group name is empty")
}

// Get 从当前分组中获取指定键的值。
// 如果没有有效的分组名，则返回空字符串。
func (g *Group) Get(key string) string {
	if g.hasGroupName() {
		return g.Cache.Get(key)
	}
	return ""
}

// Exist 检查当前分组中是否存在指定的键。
// 如果没有有效的分组名，则始终返回false。
func (g *Group) Exist(key string) bool {
	if g.hasGroupName() {
		return g.Cache.Exist(key)
	}
	return false
}

// Expire 获取当前分组中指定键的过期时间。
// 如果没有有效的分组名，则返回0。
func (g *Group) Expire(key string) time.Duration {
	if g.hasGroupName() {
		return g.Cache.Expire(key)
	}
	return time.Duration(0)
}

// Delete 从当前分组中删除指定的键。
// 如果没有有效的分组名，则操作失败。
func (g *Group) Delete(key string) error {
	if g.hasGroupName() {
		return g.Cache.Delete(key)
	}
	return errors.New("group name is empty")
}

// Clear 清除当前分组中的所有键值对。
// 如果没有有效的分组名，则操作失败。
func (g *Group) Clear() error {
	if g.hasGroupName() {
		return g.Cache.Clear()
	}
	return errors.New("group name is empty")
}
