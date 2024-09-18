package boot

import (
	`fmt`
	`io`
	`os`
	`strings`
	
	`github.com/chaodoing/boot/cache`
	`github.com/chaodoing/boot/config`
	`github.com/chaodoing/boot/container`
	`github.com/chaodoing/boot/o`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/middleware/logger`
	`github.com/kataras/iris/v12/middleware/recover`
	`gorm.io/gorm`
)

type (
	// Launch 结构体代表了一个启动配置或者状态的实例。
	Launch struct {
		app        *iris.Application
		containers container.Container
		config     iris.Configuration
		env        config.Config
		db         *gorm.DB
		cache      *cache.Cache
		group      *cache.Group
	}
	// Handle 是一个配置路由处理器的函数，负责初始化应用程序的路由和相关处理逻辑。
	// 该函数主要目的是将不同的服务组件（如containers, db, jwt）整合到app中。
	// 这里的注释解释了函数的作用、参数和没有返回值的原因。
	//
	// 参数:
	// - app: *iris.Application 类型，代表Iris web框架的应用实例，用于注册路由和处理函数。
	// - containers: container.Container 类型，是一个依赖注入容器，包含应用中需要管理的各种服务实例。
	// - db: *gorm.DB 类型，代表数据库连接对象，用于数据库操作。
	// - jwt: *container.Jwt 类型，是JWT认证相关的配置和处理逻辑，确保请求的合法性。
	// 注意，该函数没有返回值，因为它直接修改app实例的状态，而不是通过返回值来表达结果。
	Handle func(app *iris.Application, containers container.Container, db *gorm.DB, jwt *container.Jwt)
)

// New 创建一个新的应用启动配置
// 参数:
//   file - 配置文件的路径
// 返回值:
//   Launch - 应用启动实例，包含了应用和相关配置信息
func New(file string) Launch {
	// 检查测试环境是否可用
	if err := config.TestENV(); err != nil {
		panic(err)
	}
	
	// 初始化环境配置
	var (
		env config.Config
		err error
	)
	
	// 解析配置文件路径中的环境变量
	file = os.ExpandEnv(file)
	
	// 根据配置文件的扩展名解析配置
	exts := strings.Split(file, ".")
	switch exts[len(exts)-1] {
	case "json":
		env, err = config.Json(file)
	case "yaml", "yml":
		env, err = config.Yaml(file)
	case "xml":
		env, err = config.Xml(file)
	}
	
	// 如果配置加载失败，则抛出异常
	if err != nil {
		panic(err)
	}
	
	// 加载环境变量到配置中
	env = env.LoadEnv()
	
	// 初始化容器
	dock := container.New(env)
	
	// 获取数据库实例
	db, err := dock.Database()
	if err != nil {
		panic(err)
	}
	
	// 获取缓存实例
	caching, err := dock.Cache()
	if err != nil {
		panic(err)
	}
	
	// 获取分组实例
	group, err := dock.Group()
	if err != nil {
		panic(err)
	}
	
	// 注册容器内的服务
	hero.Register(dock)
	
	// 初始化 Iris 框架
	app := iris.New()
	
	// 使用全局中间件
	app.UseGlobal(iris.Compression)
	app.UseRouter(recover.New())
	app.UseRouter(logger.New())
	
	// 注册依赖项
	app.RegisterDependency(dock, db, caching, group)
	
	// 处理路由
	o.Handle(app)
	
	// 配置跨域
	if env.Service.Cross {
		app.AllowMethods(iris.MethodOptions)
		app.UseRouter(container.Cors)
	}
	
	// 配置日志输出
	var writer io.Writer
	writer, err = env.Service.Logger.Writer()
	if err != nil {
		panic(err)
	}
	app.Logger().SetOutput(writer).SetLevel(env.Service.Logger.IrisLevel())
	
	// 返回应用启动实例
	return Launch{
		app:        app,
		containers: dock,
		env:        env,
		db:         db,
		cache:      caching,
		group:      group,
	}
}

// IrisConfiguration 设置或更新 Launch 结构体中的配置信息。
// 该方法允许在创建 Iris 应用实例时，通过链式调用设置不同的配置。
// 参数:
//   config - 一个 iris.Configuration 类型的配置对象，包含了应用运行所需的各种配置。
// 返回值:
//   Launch - 返回更新后的 Launch 实例，支持链式调用。
func (l Launch) IrisConfiguration(config iris.Configuration) Launch {
	l.config = config
	return l
}

// SetEvents 是 Launch 类的一个方法，用于为当前的 Launch 实例设置事件监听器。
// 该方法接收一个类型为 map[string]interface{} 的参数 e，其中包含了事件名和对应的处理方法。
// 返回值为修改后的 Launch 实例，允许链式调用。
func (l Launch) SetEvents(e map[string]interface{}) Launch {
	// 遍历参数 e 中的所有键值对，即事件名和处理方法。
	for name, method := range e {
		// 调用 l.containers.Events 的 AddEventListener 方法，为每个事件名添加对应的处理方法。
		// 这里实际进行的是事件监听器的注册过程。
		l.containers.Events.AddEventListener(name, method)
	}
	// 返回修改后的 Launch 实例，支持链式调用。
	return l
}

// DB 方法返回与 Launch 实例关联的数据库连接。
// 该方法主要用于获取 Launch 结构体内的数据库连接实例，以便在其他地方操作数据库。
func (l Launch) DB() *gorm.DB {
	return l.db
}

// Launch 的 Handle 方法接收一个或多个Handle函数作为参数，并依次调用这些函数，
// 每个函数都会执行一些处理逻辑。这种方法常用于在启动过程中初始化或配置各个组件。
// 此设计模式允许在保持代码灵活和可扩展的同时，有序地执行初始化逻辑。
//
// 参数:
//   values ...Handle: 一个或多个Handle类型的函数，这些函数会在方法执行时被调用。
// 返回值:
//   Launch: 方法执行完毕后返回Launch实例本身，支持链式调用。
func (l Launch) Handle(values ...Handle) Launch {
	// 遍历传入的Handle函数切片
	for _, fn := range values {
		// 调用每一个函数，并传递Launch实例的app、containers、db以及JWT令牌给它们
		// 这允许每个函数对这些资源执行特定的初始化或处理逻辑
		fn(l.app, l.containers, l.db, l.containers.Jwt())
	}
	// 返回Launch实例本身，支持链式调用
	return l
}

// Run 方法用于启动服务。
// 该方法主要负责配置服务、设置路由并启动服务。
func (l Launch) Run() {
	// 设置PostMaxMemory，以支持的最大内存大小乘以MB（每MB为1024KB）。
	l.config.PostMaxMemory = l.env.Uploader.MaxSize * iris.MB
	// 根据环境变量"ENV"的值来决定是否禁用启动日志。
	// 如果环境为"development"，则不禁用启动日志。
	l.config.DisableStartupLog = !strings.EqualFold(os.Getenv("ENV"), "development")
	// 设置其他配置，当前主要是设置路由信息。
	// 使用map存储其他配置项，其中"routes"键对应服务的所有路由。
	l.config.Other = map[string]interface{}{
		"routes": l.app.GetRoutes(),
	}
	// 启动服务，使用配置好的地址和端口，以及自定义配置。
	// 如果启动过程中出现错误，则抛出异常。
	err := l.app.Run(iris.Addr(fmt.Sprintf("%s:%d", l.env.Service.Host, l.env.Service.Port)), iris.WithConfiguration(l.config))
	if err != nil {
		panic(err)
	}
}
