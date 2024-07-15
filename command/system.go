package command

import (
	`bytes`
	`os`
	`text/template`
	
	`github.com/chaodoing/boot/assets/service`
	`github.com/chaodoing/boot/config`
	`github.com/gookit/goutil/envutil`
	`github.com/urfave/cli`
)

var System = cli.Command{
	Name:        "system",
	Description: "生成Linux服务脚本",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "exec",
			Usage:    "服务运行命令",
			Required: true,
		},
	},
	Action: func(c *cli.Context) (err error) {
		err = config.TestENV()
		if err != nil {
			return
		}
		tpl, err := template.New("systemd").Parse(service.Systemd)
		if err != nil {
			return
		}
		buf := new(bytes.Buffer)
		err = tpl.Execute(buf, map[string]string{
			"description": "应用描述内容",
			"username":    "root",
			"group":       "root",
			"directory":   os.Getenv("WORKDIR"),
			"configDir":   os.Getenv("CONFIG_DIR"),
			"logDir":      os.Getenv("LOG_DIR"),
			"env":         os.Getenv("ENV"),
			"version":     envutil.Getenv("VERSION", "v1.0.0"),
			"execute":     c.String("exec"),
		})
		_, err = os.Stdout.Write(buf.Bytes())
		return
	},
}
