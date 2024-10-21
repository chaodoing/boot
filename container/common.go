package container

import (
	`github.com/chaodoing/boot/config`
	`github.com/chaodoing/boot/task`
)

func New(config config.Config) Container {
	return Container{
		Config:  config,
		Crontab: task.NewCrontab(),
		Events:  &task.Events{},
	}
}
