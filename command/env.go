package command

import (
	`github.com/chaodoing/boot/config`
	`github.com/urfave/cli`
)

var Env = cli.Command{
	Name:        "env",
	Description: "生成环境变量文件",
	Category:    "Frame",
	Action: func(c *cli.Context) error {
		var env = config.Default()
		err := config.INIWriter(env)
		if err != nil {
			return err
		}
		return nil
	},
}
