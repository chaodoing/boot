package orm

import (
	`fmt`
	`os`
	
	`gopkg.in/ini.v1`
)

func Environment() (c *ini.File, e error) {
	env := os.ExpandEnv(fmt.Sprintf("${WORKDIR}/%s", "env"))
	development := os.ExpandEnv(fmt.Sprintf("${WORKDIR}/%s.development", env))
	production := os.ExpandEnv(fmt.Sprintf("${WORKDIR}/%s.production", env))
	return ini.LoadSources(ini.LoadOptions{
		AllowShadows:             true,
		IgnoreInlineComment:      true,
		SpaceBeforeInlineComment: true,
		AllowBooleanKeys:         true,
	}, env, development, production)
}
