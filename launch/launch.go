package launch

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
	Launch struct {
		app        *iris.Application
		containers container.Container
		config     iris.Configuration
		env        config.Config
		db         *gorm.DB
		cache      *cache.Cache
		group      *cache.Group
	}
	Handle func(app *iris.Application, containers container.Container)
)

func New(file string) Launch {
	if err := config.TestENV(); err != nil {
		panic(err)
	}
	var (
		env config.Config
		err error
	)
	file = os.ExpandEnv(file)
	ss := strings.Split(file, ".")
	switch ss[len(ss)-1] {
	case "json":
		env, err = config.Json(file)
	case "yaml", "yml":
		env, err = config.Yaml(file)
	case "xml":
		env, err = config.Xml(file)
	}
	if err != nil {
		panic(err)
	}
	env = env.LoadEnv()
	
	dock := container.New(env)
	db, err := dock.Database()
	if err != nil {
		panic(err)
	}
	caching, err := dock.Cache()
	if err != nil {
		panic(err)
	}
	group, err := dock.Group()
	if err != nil {
		panic(err)
	}
	hero.Register(dock)
	app := iris.New()
	app.UseGlobal(iris.Compression)
	app.UseRouter(recover.New())
	app.UseRouter(logger.New())
	app.RegisterDependency(dock, db, caching, group)
	o.Handle(app)
	if env.Service.Cross {
		app.AllowMethods(iris.MethodOptions)
		app.UseGlobal(container.Cors)
	}
	var writer io.Writer
	writer, err = env.Service.Logger.Writer()
	if err != nil {
		panic(err)
	}
	app.Logger().SetOutput(writer).SetLevel(env.Service.Logger.IrisLevel())
	return Launch{
		app:        app,
		containers: dock,
		env:        env,
		db:         db,
		cache:      caching,
		group:      group,
	}
}
func (l Launch) IrisConfiguration(config iris.Configuration) Launch {
	l.config = config
	return l
}

func (l Launch) DB() *gorm.DB {
	return l.db
}

func (l Launch) Handle(values ...Handle) Launch {
	for _, fn := range values {
		fn(l.app, l.containers)
	}
	return l
}
func (l Launch) Run() {
	l.config.PostMaxMemory = l.env.Uploader.MaxSize * iris.MB
	l.config.DisableStartupLog = !strings.EqualFold(os.Getenv("ENV"), "development")
	l.config.Other = map[string]interface{}{
		"routes": l.app.GetRoutes(),
	}
	err := l.app.Run(iris.Addr(fmt.Sprintf("%s:%d", l.env.Service.Host, l.env.Service.Port)), iris.WithConfiguration(l.config))
	if err != nil {
		panic(err)
	}
}
