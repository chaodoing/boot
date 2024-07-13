package models

import (
	"github.com/chaodoing/boot/calendar"
	"gorm.io/datatypes"
)

// AccessLogs 操作日志
type AccessLogs struct {
	ID      uint64            `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null;comment:'主键'" json:"id"` // 主键
	UID     uint32            `gorm:"column:uid;type:int(10) unsigned;not null;default:0;comment:'用户id'" json:"uid"`                    // 用户id
	Node    uint8             `gorm:"column:node;type:tinyint(3) unsigned;not null;default:0;comment:'所属后台'" json:"node"`               // 所属后台
	Method  string            `gorm:"column:method;type:enum('GET','POST','PUT','DELETE');not null;comment:'请求方法'" json:"method"`       // 请求方法
	Action  string            `gorm:"column:action;type:varchar(255);not null;comment:'操作地址'" json:"action"`                            // 操作地址
	Query   string            `gorm:"column:query;type:varchar(255);not null;comment:'访问参数'" json:"query"`                              // 访问参数
	Body    string            `gorm:"column:body;type:text;default:null;comment:'请求内容'" json:"body"`                                    // 请求内容
	Address string            `gorm:"column:address;type:varchar(30);not null;comment:'请求的IP'" json:"address"`                          // 请求的IP
	TimeAt  calendar.Datetime `gorm:"column:time_at;type:datetime;not null;comment:'创建时间'" json:"timeAt"`                               // 创建时间
}

// TableName get sql table name.获取数据库表名
func (m *AccessLogs) TableName() string {
	return "access_logs"
}

// Administrator 管理员信息
type Administrator struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	RoleID   uint32            `gorm:"column:role_id;type:int(10) unsigned;not null;comment:'所属角色'" json:"roleId"`                    // 所属角色
	Avatar   string            `gorm:"column:avatar;type:varchar(255);not null;comment:'用户头像'" json:"avatar"`                         // 用户头像
	FullName string            `gorm:"column:full_name;type:varchar(20);not null;comment:'姓名'" json:"fullName"`                       // 姓名
	Username string            `gorm:"column:username;type:varchar(60);not null;comment:'登录账号'" json:"username"`                      // 登录账号
	Password string            `gorm:"column:password;type:varchar(50);not null;comment:'登录密码'" json:"password"`                      // 登录密码
	Gender   uint8             `gorm:"column:gender;type:tinyint(3) unsigned;not null;default:0;comment:'用户性别'" json:"gender"`        // 用户性别
	Locked   int8              `gorm:"column:locked;type:tinyint(4);not null;comment:'锁定状态'" json:"locked"`                           // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Administrator) TableName() string {
	return "administrator"
}

// Article 文章信息表
type Article struct {
	ID             uint64            `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint(20) unsigned;not null;comment:'主键'" json:"id"`           // 主键
	ParentID       uint64            `gorm:"column:parent_id;type:bigint(20) unsigned;not null;default:0;comment:'上级文章id'" json:"parentId"`              // 上级文章id
	Type           uint8             `gorm:"column:type;type:tinyint(3) unsigned;not null;default:1;comment:'文章类型 [1:文章|2:页面]'" json:"type"`             // 文章类型 [1:文章|2:页面]
	Title          string            `gorm:"column:title;type:varchar(255);not null;comment:'标题'" json:"title"`                                          // 标题
	SeoKeywords    string            `gorm:"column:seo_keywords;type:varchar(255);not null;default:'';comment:'页面关键字'" json:"seoKeywords"`               // 页面关键字
	SeoDescription string            `gorm:"column:seo_description;type:varchar(255);not null;default:'';comment:'文章描述'" json:"seoDescription"`          // 文章描述
	Content        string            `gorm:"column:content;type:text;not null;comment:'文章内容'" json:"content"`                                            // 文章内容
	Thumbnail      string            `gorm:"column:thumbnail;type:varchar(255);not null;default:'';comment:'缩略图'" json:"thumbnail"`                      // 缩略图
	AdminID        uint32            `gorm:"column:admin_id;type:int(10) unsigned;not null;default:0;comment:'发表者用户ID'" json:"adminId"`                  // 发表者用户ID
	Status         uint8             `gorm:"column:status;type:tinyint(3) unsigned;not null;default:0;comment:'发布状态 [0:未发布|1:已发布]'" json:"status"`       // 发布状态 [0:未发布|1:已发布]
	Format         uint8             `gorm:"column:format;type:tinyint(3) unsigned;not null;default:1;comment:'内容格式 [1:html|2:markdown]'" json:"format"` // 内容格式 [1:html|2:markdown]
	Recommended    uint8             `gorm:"column:recommended;type:tinyint(3) unsigned;not null;default:0;comment:'是否推荐'" json:"recommended"`           // 是否推荐
	IsTop          uint8             `gorm:"column:is_top;type:tinyint(3) unsigned;not null;default:0;comment:'是否置顶'" json:"isTop"`                      // 是否置顶
	Favorites      uint32            `gorm:"column:favorites;type:int(10) unsigned;not null;default:0;comment:'收藏数'" json:"favorites"`                   // 收藏数
	Hits           uint64            `gorm:"column:hits;type:bigint(20) unsigned;not null;default:0;comment:'点击数'" json:"hits"`                          // 点击数
	Like           uint32            `gorm:"column:like;type:int(10) unsigned;not null;default:0;comment:'点赞数'" json:"like"`                             // 点赞数
	Source         string            `gorm:"column:source;type:varchar(255);not null;default:'';comment:'转载文章的来源'" json:"source"`                        // 转载文章的来源
	CreateAt       calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                                     // 创建时间
	PublishedAt    calendar.Datetime `gorm:"column:published_at;type:datetime;default:null;comment:'发布时间'" json:"publishedAt"`                           // 发布时间
	UpdateAt       calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                                 // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "article"
}

// Column 文章栏目
type Column struct {
	ID             uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"`      // 主键
	Code           string            `gorm:"column:code;type:varchar(255);not null;comment:'分类编码'" json:"code"`                                  // 分类编码
	ParentID       uint32            `gorm:"column:parent_id;type:int(10) unsigned;not null;comment:'上级id'" json:"parentId"`                     // 上级id
	Thumbnail      string            `gorm:"column:thumbnail;type:varchar(255);not null;default:'';comment:'缩略图'" json:"thumbnail"`              // 缩略图
	PublishCount   uint32            `gorm:"column:publish_count;type:int(10) unsigned;not null;default:0;comment:'发布文章数量'" json:"publishCount"` // 发布文章数量
	Title          string            `gorm:"column:title;type:varchar(255);not null;comment:'分类名称'" json:"title"`                                // 分类名称
	Description    string            `gorm:"column:description;type:varchar(255);not null;default:'';comment:'分类描述'" json:"description"`         // 分类描述
	Content        string            `gorm:"column:content;type:text;not null;comment:'分类内容'" json:"content"`                                    // 分类内容
	SeoTitle       string            `gorm:"column:seo_title;type:varchar(255);not null;default:'';comment:'网页标题'" json:"seoTitle"`              // 网页标题
	SeoKeywords    string            `gorm:"column:seo_keywords;type:varchar(255);not null;default:'';comment:'网页关键字'" json:"seoKeywords"`       // 网页关键字
	SeoDescription string            `gorm:"column:seo_description;type:varchar(255);not null;default:'';comment:'网页描述'" json:"seoDescription"`  // 网页描述
	ListSort       uint32            `gorm:"column:list_sort;type:int(10) unsigned;not null;comment:'排序'" json:"listSort"`                       // 排序
	Locked         uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`             // 锁定状态
	CreateAt       calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                             // 创建时间
	UpdateAt       calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                         // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Column) TableName() string {
	return "column"
}

// Config 配置信息表
type Config struct {
	ID       int16             `gorm:"autoIncrement:true;primaryKey;column:id;type:mediumint(8) unsigned;not null;comment:'主键'" json:"id"` // 主键
	Code     string            `gorm:"column:code;type:varchar(60);not null;comment:'调用编码'" json:"code"`                                   // 调用编码
	Title    string            `gorm:"column:title;type:varchar(30);not null;comment:'配置名称'" json:"title"`                                 // 配置名称
	Node     uint8             `gorm:"column:node;type:tinyint(3) unsigned;not null;default:0;comment:'所属后台节点'" json:"node"`               // 所属后台节点
	Sort     int16             `gorm:"column:sort;type:mediumint(8) unsigned;not null;comment:'显示顺序'" json:"sort"`                         // 显示顺序
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`             // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                             // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                         // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Config) TableName() string {
	return "config"
}

// ConfigValue 配置内容信息
type ConfigValue struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	ConfigID int16             `gorm:"column:config_id;type:mediumint(8) unsigned;not null;comment:'配置ID'" json:"configId"`           // 配置ID
	Title    string            `gorm:"column:title;type:varchar(50);not null;comment:'配置标题'" json:"title"`                            // 配置标题
	Code     string            `gorm:"column:code;type:varchar(60);not null;comment:'配置编码'" json:"code"`                              // 配置编码
	Value    string            `gorm:"column:value;type:text;not null;comment:'配置值'" json:"value"`                                    // 配置值
	Options  datatypes.JSON    `gorm:"column:options;type:json;not null;comment:'配置项'" json:"options"`                                // 配置项
	Type     string            `gorm:"column:type;type:varchar(32);not null;comment:'配置类型'" json:"type"`                              // 配置类型
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'排序'" json:"sort"`                           // 排序
	Note     string            `gorm:"column:note;type:varchar(120);not null;comment:'配置说明'" json:"note"`                             // 配置说明
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`        // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *ConfigValue) TableName() string {
	return "config_value"
}

// Developer 开发后台用户
type Developer struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	RoleID   uint32            `gorm:"column:role_id;type:int(10) unsigned;not null;comment:'用户角色'" json:"roleId"`                    // 用户角色
	Avatar   string            `gorm:"column:avatar;type:varchar(255);not null;comment:'用户头像'" json:"avatar"`                         // 用户头像
	FullName string            `gorm:"column:full_name;type:varchar(12);not null;comment:'用户姓名'" json:"fullName"`                     // 用户姓名
	Username string            `gorm:"column:username;type:varchar(50);not null;comment:'登录账号'" json:"username"`                      // 登录账号
	Password string            `gorm:"column:password;type:varchar(50);not null;comment:'登录密码'" json:"password"`                      // 登录密码
	Gender   uint8             `gorm:"column:gender;type:tinyint(3) unsigned;not null;default:0;comment:'用户性别'" json:"gender"`        // 用户性别
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`        // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Developer) TableName() string {
	return "developer"
}

// Dict 字典类型表
type Dict struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	Title    string            `gorm:"column:title;type:varchar(50);not null;comment:'字典名称'" json:"title"`                            // 字典名称
	Code     string            `gorm:"column:code;type:varchar(50);not null;comment:'字典值'" json:"code"`                               // 字典值
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'显示顺序'" json:"sort"`                         // 显示顺序
	Note     string            `gorm:"column:note;type:varchar(255);not null;comment:'字典备注'" json:"note"`                             // 字典备注
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Dict) TableName() string {
	return "dict"
}

// DictData 字典项信息表
type DictData struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	DictID   uint32            `gorm:"column:dict_id;type:int(10) unsigned;not null;comment:'字典类型id'" json:"dictId"`                  // 字典类型id
	Title    string            `gorm:"column:title;type:varchar(50);not null;comment:'字典名称'" json:"title"`                            // 字典名称
	Name     string            `gorm:"column:name;type:varchar(50);not null;comment:'字典编码'" json:"name"`                              // 字典编码
	Value    string            `gorm:"column:value;type:text;not null;comment:'字典内容'" json:"value"`                                   // 字典内容
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'显示顺序'" json:"sort"`                         // 显示顺序
	Locked   int8              `gorm:"column:locked;type:tinyint(4);default:null;comment:'锁定状态'" json:"locked"`                       // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *DictData) TableName() string {
	return "dict_data"
}

// Navigate 菜单信息表
type Navigate struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键id'" json:"id"`  // 主键id
	Pid      uint32            `gorm:"column:pid;type:int(10) unsigned;not null;default:0;comment:'上级菜单'" json:"pid"`                    // 上级菜单
	Title    string            `gorm:"column:title;type:varchar(30);not null;comment:'菜单标题'" json:"title"`                               // 菜单标题
	Code     string            `gorm:"unique;column:code;type:varchar(60);not null;comment:'唯一名称 可用于权限检查'" json:"code"`                  // 唯一名称 可用于权限检查
	Icon     string            `gorm:"column:icon;type:varchar(255);not null;comment:'菜单图标'" json:"icon"`                                // 菜单图标
	Href     string            `gorm:"column:href;type:varchar(120);not null;comment:'URL地址'" json:"href"`                               // URL地址
	Query    string            `gorm:"column:query;type:varchar(120);not null;default:'';comment:'请求参数'" json:"query"`                   // 请求参数
	Target   string            `gorm:"column:target;type:enum('_self','_blank','_parent','_top');not null;comment:'打开方式'" json:"target"` // 打开方式
	Mold     uint8             `gorm:"column:mold;type:tinyint(3) unsigned;not null;default:1;comment:'类型'" json:"mold"`                 // 类型
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'菜单排序'" json:"sort"`                            // 菜单排序
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`           // 锁定状态
	Remark   string            `gorm:"column:remark;type:varchar(255);not null;default:'';comment:'菜单备注'" json:"remark"`                 // 菜单备注
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                           // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                       // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Navigate) TableName() string {
	return "navigate"
}

// Position 广告位信息表
type Position struct {
	ID          uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	Title       string            `gorm:"column:title;type:varchar(60);not null;comment:'广告位名称'" json:"title"`                           // 广告位名称
	Code        string            `gorm:"column:code;type:varchar(60);not null;comment:'广告位编码'" json:"code"`                             // 广告位编码
	Description string            `gorm:"column:description;type:varchar(255);not null;comment:'广告位描述'" json:"description"`              // 广告位描述
	Width       int16             `gorm:"column:width;type:mediumint(8) unsigned;not null;comment:'广告位图片宽度'" json:"width"`               // 广告位图片宽度
	Height      int16             `gorm:"column:height;type:mediumint(8) unsigned;not null;comment:'广告位图片高度'" json:"height"`             // 广告位图片高度
	Sort        int16             `gorm:"column:sort;type:mediumint(8) unsigned;not null;comment:'广告位排序'" json:"sort"`                   // 广告位排序
	Locked      uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`        // 锁定状态
	CreateAt    calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt    calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Position) TableName() string {
	return "position"
}

// Role 角色信息表
type Role struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"` // 主键
	Node     uint8             `gorm:"column:node;type:tinyint(3) unsigned;not null;default:0;comment:'所属后台'" json:"node"`            // 所属后台
	Title    string            `gorm:"column:title;type:varchar(60);not null;comment:'角色名称'" json:"title"`                            // 角色名称
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'排序'" json:"sort"`                           // 排序
	Rules    datatypes.JSON    `gorm:"column:rules;type:json;not null;comment:'角色允许访问'" json:"rules"`                                 // 角色允许访问
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;comment:'角色状态'" json:"locked"`                  // 角色状态
	Note     string            `gorm:"column:note;type:varchar(120);not null;comment:'角色说明'" json:"note"`                             // 角色说明
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                        // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                    // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Role) TableName() string {
	return "role"
}

// Rule 菜单信息表
type Rule struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键id'" json:"id"`  // 主键id
	Pid      uint32            `gorm:"column:pid;type:int(10) unsigned;not null;default:0;comment:'上级菜单'" json:"pid"`                    // 上级菜单
	Title    string            `gorm:"column:title;type:varchar(30);not null;comment:'菜单标题'" json:"title"`                               // 菜单标题
	Code     string            `gorm:"column:code;type:varchar(60);not null;comment:'唯一名称 可用于权限检查'" json:"code"`                         // 唯一名称 可用于权限检查
	Icon     string            `gorm:"column:icon;type:varchar(60);not null;comment:'菜单图标'" json:"icon"`                                 // 菜单图标
	Href     string            `gorm:"column:href;type:varchar(120);not null;comment:'URL地址'" json:"href"`                               // URL地址
	Query    string            `gorm:"column:query;type:varchar(120);not null;default:'';comment:'请求参数'" json:"query"`                   // 请求参数
	Target   string            `gorm:"column:target;type:enum('_self','_blank','_parent','_top');not null;comment:'打开方式'" json:"target"` // 打开方式
	Node     int8              `gorm:"column:node;type:tinyint(4);not null;default:0;comment:'所属后台节点'" json:"node"`                      // 所属后台节点
	Mold     uint8             `gorm:"column:mold;type:tinyint(3) unsigned;not null;default:1;comment:'类型'" json:"mold"`                 // 类型
	Sort     uint32            `gorm:"column:sort;type:int(10) unsigned;not null;comment:'菜单排序'" json:"sort"`                            // 菜单排序
	Hide     uint8             `gorm:"column:hide;type:tinyint(3) unsigned;not null;default:0;comment:'是否隐藏'" json:"hide"`               // 是否隐藏
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`           // 锁定状态
	Remark   string            `gorm:"column:remark;type:varchar(255);not null;default:'';comment:'菜单备注'" json:"remark"`                 // 菜单备注
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                           // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                       // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Rule) TableName() string {
	return "rule"
}

// Slider 轮播图信息表
type Slider struct {
	ID       uint32            `gorm:"autoIncrement:true;primaryKey;column:id;type:int(10) unsigned;not null;comment:'主键'" json:"id"`    // 主键
	Pid      uint32            `gorm:"column:pid;type:int(10) unsigned;not null;comment:'所属位置'" json:"pid"`                              // 所属位置
	Title    string            `gorm:"column:title;type:varchar(100);not null;comment:'轮播标题'" json:"title"`                              // 轮播标题
	Href     string            `gorm:"column:href;type:varchar(120);not null;comment:'轮播图跳转地址'" json:"href"`                             // 轮播图跳转地址
	Target   string            `gorm:"column:target;type:enum('_self','_blank','_parent','_top');not null;comment:'打开方式'" json:"target"` // 打开方式
	Photo    string            `gorm:"column:photo;type:varchar(255);not null;comment:'所属图片地址'" json:"photo"`                            // 所属图片地址
	Sort     int16             `gorm:"column:sort;type:mediumint(8) unsigned;not null;comment:'图片排序'" json:"sort"`                       // 图片排序
	Locked   uint8             `gorm:"column:locked;type:tinyint(3) unsigned;not null;default:0;comment:'锁定状态'" json:"locked"`           // 锁定状态
	CreateAt calendar.Datetime `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"createAt"`                           // 创建时间
	UpdateAt calendar.Datetime `gorm:"column:update_at;type:datetime;default:null;comment:'修改时间'" json:"updateAt"`                       // 修改时间
}

// TableName get sql table name.获取数据库表名
func (m *Slider) TableName() string {
	return "slider"
}
