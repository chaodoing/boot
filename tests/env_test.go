package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/config`
	`github.com/gookit/goutil/envutil`
)

func TestEnvironment(t *testing.T) {
	envutil.SetEnvMap(map[string]string{
		"WORKDIR": "/Users/superman/Server/src/github.com/chaodoing/boot",
		"ENV":     "development",
	})
	cc := config.Config{}
	t.Error(cc.LoadEnv())
	// conf, err := ini.Load(os.ExpandEnv("${WORKDIR}/.env.${ENV}"))
	// if err != nil {
	// 	t.Error(err)
	// }
	// if err := conf.MapTo(&cc); err != nil {
	// 	t.Error(err)
	// }
	// t.Log(cc)
}
