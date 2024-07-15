package command

import (
	`encoding/xml`
	`os`
	
	`github.com/chaodoing/boot/config`
	`github.com/gookit/goutil/fsutil`
	`github.com/urfave/cli`
)

var Config = cli.Command{
	Name:        "config",
	Description: "生成配置文件",
	Category:    "Frame",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "export",
			Usage:    "配置文件名称",
			Required: true,
			Value:    os.ExpandEnv("${CONFIG_DIR}"),
		},
	},
	Action: func(ctx *cli.Context) error {
		var env = config.Default()
		value, err := xml.MarshalIndent(env, "", "\t")
		if err != nil {
			return err
		}
		value = append([]byte(xml.Header), value...)
		_, err = fsutil.PutContents(ctx.String("export"), value)
		if err != nil {
			return err
		}
		return nil
	},
}
