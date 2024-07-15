package service

import (
	_ `embed`
)

//go:embed systemd.service
var Systemd string
