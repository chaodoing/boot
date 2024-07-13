package tests

import (
	`testing`
	
	`github.com/chaodoing/boot/config`
	`github.com/gookit/goutil/envutil`
)

func TestConfig(t *testing.T) {
	envutil.SetEnvMap(map[string]string{
		"WORKDIR": "/Users/superman/Server/src/github.com/chaodoing/boot",
		"ENV":     "development",
		"VERSION": "v0.0.1",
	})
	var env = config.Default()
	err := config.INIWriter(env)
	if err != nil {
		t.Error(err)
	}
	t.Log("Success")
}
