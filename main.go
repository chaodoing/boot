package main

import (
	`os`
	
	`github.com/chaodoing/boot/container`
	`github.com/chaodoing/boot/http/common`
	`github.com/chaodoing/boot/http/controllers`
	`github.com/chaodoing/boot/launch`
	`github.com/chaodoing/boot/models`
	`github.com/chaodoing/boot/traits`
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/mvc`
)

func init() {
	envutil.SetEnvMap(map[string]string{
		"WORKDIR":    os.ExpandEnv("${PWD}"),
		"CONFIG_DIR": os.ExpandEnv("${PWD}/custom"),
		"LOG_DIR":    os.ExpandEnv("${PWD}/logs"),
		"ENV":        "development",
		"VERSION":    "v1.0.0",
	})
}

func main() {
	boot := launch.New("${CONFIG_DIR}/config.xml")
	var c launch.Handle = func(app *iris.Application, box container.Container) {
		index := app.Party("/")
		{
			var m = models.ConfigValue{
				Model: &traits.Model{
					DB: boot.DB(),
				},
			}
			mvc.New(index).Register(m).Handle(new(controllers.Index)).SetName("控制器")
		}
		app.Get("/index.html", hero.Handler(common.Home)).Name = "测试"
	}
	boot.Handle(c).Run()
}
