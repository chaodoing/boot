package main

import (
	`os`
	
	`github.com/chaodoing/boot/container`
	`github.com/chaodoing/boot/http/common`
	`github.com/chaodoing/boot/http/controllers`
	`github.com/chaodoing/boot/launch`
	`github.com/gookit/goutil/envutil`
	`github.com/kataras/iris/v12`
	`github.com/kataras/iris/v12/hero`
	`github.com/kataras/iris/v12/mvc`
)

func main() {
	envutil.SetEnvMap(map[string]string{
		"WORKDIR":    os.ExpandEnv("${PWD}"),
		"CONFIG_DIR": os.ExpandEnv("${PWD}/custom"),
		"LOG_DIR":    os.ExpandEnv("${PWD}/logs"),
		"ENV":        "development",
		"VERSION":    "v1.0.0",
	})
	boot := launch.New("${CONFIG_DIR}/config.xml")
	var c launch.Handle = func(app *iris.Application, box container.Container) {
		index := app.Party("/")
		{
			mvc.New(index).Register(box).Handle(new(controllers.Index))
		}
		app.Get("/index.html", hero.Handler(common.Home)).Name = "测试"
	}
	boot.Handle(c).Run()
}
