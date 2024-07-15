package command

import (
	`github.com/urfave/cli`
)

var System = cli.Command{
	Name:        "systemctl",
	Description: "生成Linux服务脚本",
	Category:    "Frame",
	Action: func(c *cli.Context) (err error) {
		
		return
	},
}
