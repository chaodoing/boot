package task

import (
	`github.com/jasonlvhit/gocron`
)

type Crontab struct {
	crontab *gocron.Scheduler
}

func NewCrontab() *Crontab {
	return &Crontab{
		crontab: gocron.NewScheduler(),
	}
}

func (c *Crontab) Every(interval uint64) *gocron.Job {
	return c.crontab.Every(interval)
}

func (c *Crontab) Clear() {
	c.crontab.Clear()
}

func (c *Crontab) Start() {
	<-c.crontab.Start()
}
