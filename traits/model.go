package traits

import (
	`github.com/kataras/iris/v12`
	`gorm.io/gorm`
)

type Model struct {
	DB      *gorm.DB `gorm:"-" json:"-" xml:"-" yaml:"-"`
	Offset  int      `gorm:"-" json:"-" xml:"-" yaml:"-"`
	Limit   int      `gorm:"-" json:"-" xml:"-" yaml:"-"`
	Total   int64    `gorm:"-" json:"-" xml:"-" yaml:"-"`
	Current int      `gorm:"-" json:"-" xml:"-" yaml:"-"`
}

// Pagination 实现分页查询。
// 使用 where 参数来指定查询条件，通过设置偏移量（Offset）和限制（Limit）来实现分页。
// 函数返回一个错误，如果在查询过程中发生了错误。
//
// 参数:
//   where ...interface{}: 查询条件，可以是任意数量的参数，用于构建 WHERE 子句。
//
// 返回值:
//   error: 查询过程中可能发生的错误。
func (m Model) Pagination(ctx iris.Context, table string) (tx *gorm.DB, err error) {
	// 从URL参数中尝试获取页码，优先使用page参数，如果未提供，则使用current参数，两者都未提供则默认为1。
	m.Current = ctx.URLParamIntDefault("page", ctx.URLParamIntDefault("current", 1))
	// 从URL参数中尝试获取每页的项数，优先使用limit参数，如果未提供，则使用size参数，两者都未提供则默认为15。
	m.Limit = ctx.URLParamIntDefault("limit", ctx.URLParamIntDefault("size", 15))
	// 计算偏移量，用于数据库查询的OFFSET子句，确保页码从1开始。
	m.Offset = (m.Current - 1) * m.Limit
	
	// 统计满足条件的总记录数，用于后续的分页计算。
	err = m.DB.Table(table).Count(&m.Total).Error
	if err != nil {
		return
	}
	// 根据当前的偏移量和限制，查询满足条件的具体数据。
	return m.DB.Table(table).Offset(m.Offset).Limit(m.Limit), nil
}

// Field 通过指定字段进行查询
//
// 该方法用于在数据库查询中仅选择特定的字段。这可以减少不必要的数据传输，
// 提高查询效率，特别是当只需要部分字段时。
//
// 参数:
//   fields []string - 需要查询的字段列表
//
// 返回值:
//   *gorm.DB - 返回一个指向gorm.DB对象的指针，允许进一步的链式调用
func (m Model) Field(fields []string) *gorm.DB {
	// 使用Select方法指定查询的字段，并返回相应的数据库查询对象
	return m.DB.Select(fields)
}

// Where 根据给定的查询条件和参数，扩展数据库查询。
// 这个方法允许在现有的模型查询上添加WHERE子句。
// 参数query可以是一个字符串表达式，一个结构体，或者一个包含查询条件的map。
// 参数args是用于查询替换的参数，它们的顺序对应于query中的占位符。
// 返回值是一个指向gorm.DB的指针，允许进一步的数据库操作。
func (m Model) Where(query interface{}, args ...interface{}) Model {
	m.DB = m.DB.Where(query, args...)
	return m
}

// Delete 删除符合条件的数据库记录。
// 这个方法用于标记模型实例为删除状态，并在数据库中执行删除操作。
// 可选的conditions参数可以用于指定删除的条件，如果没有提供，则删除指定的模型实例。
// 返回值是一个指向gorm.DB的指针，允许进一步的数据库操作。
func (m Model) Delete(conditions ...interface{}) *gorm.DB {
	return m.DB.Delete(&m, conditions...)
}

// Update 使用提供的数据更新模型实例。
// 这个方法封装了 gorm 的 Updates 方法，简化了更新操作的调用。
// 参数 data 是一个映射，其中键是数据库表的字段名，值是要更新的字段值。
// 返回值是 *gorm.DB 类型，可以用于进一步的错误处理或链式调用。
// 方法的实现利用了指针引用，确保更新操作直接作用于原始数据，避免了数据拷贝的开销。
func (m Model) Update(data map[string]interface{}) *gorm.DB {
	return m.DB.Updates(&data)
}

// Create 在数据库中创建一个新的记录。
//
// 本方法使用 GORM 的 Create 函数来执行创建操作。它接受一个指向当前模型实例的指针，
// 并尝试将该实例插入到数据库中。如果创建成功，将返回一个指向 GORM DB 实例的指针，
// 可以通过该实例进一步查询操作的结果，如是否成功创建、是否有错误发生等。
//
// 返回值:
// *gorm.DB - 一个指向 GORM DB 实例的指针，包含操作结果信息。
func (m Model) Create() *gorm.DB {
	return m.DB.Create(m)
}

// InsertAll 批量插入数据。
//
// 参数 data 是一个接口切片，其中每个元素代表一个要插入的实体。
// 这种设计允许插入不同类型的数据，只要它们符合数据表的结构。
//
// 返回值是一个 *gorm.DB 对象，可以用于进一步的操作，如检查错误或执行事务。
// 通过调用 CreateInBatches 方法，数据将被分批插入，批次的数量等于数据切片的长度。
// 这种分批插入的方法可以有效地减少数据库的负载，提高插入性能。
func (m Model) InsertAll(data []interface{}) *gorm.DB {
	return m.DB.CreateInBatches(&data, len(data))
}

// Save 方法用于保存当前模型实例。
// 它利用内部的 DB 实例执行保存操作，将模型的更改持久化到数据库。
// 返回值为 *gorm.DB，允许进一步的操作或错误处理。
func (m Model) Save() *gorm.DB {
	// 调用 DB 的 Save 方法，传入当前模型实例进行保存
	return m.DB.Save(m)
}

// First 根据条件查询第一条记录。
// 本方法用于从数据库中检索满足给定条件的第一条记录，并将其加载到当前的Model实例中。
// 参数 where 可以接收一个或多个条件参数，用于构建查询条件。
// 返回值为错误信息，如果查询过程中发生错误，则返回相应的错误。
func (m Model) First(where ...interface{}) error {
	// 使用Gorm的First方法查询第一条满足条件的记录，并将结果错误返回。
	return m.DB.First(&m, where...).Error
}
