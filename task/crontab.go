package task

import (
	`github.com/jasonlvhit/gocron`
)

// Crontab 定义了一个结构体，用于封装 gocron.Scheduler 类型的调度器。
// 它提供了一种方式来管理需要定期执行的任务。
type Crontab struct {
	crontab *gocron.Scheduler // crontab 字段保存了一个 gocron.Scheduler 类型的调度器实例，用于控制和管理定时任务。
}

// NewCrontab 创建并返回一个新的Crontab实例。
// 该函数初始化了一个新的Crontab结构体，并将其内部的crontab字段设置为一个新的gocron调度器。
// 返回值是Crontab类型的指针，使用指针确保在外部对Crontab实例的修改可以被保持。
func NewCrontab() *Crontab {
	return &Crontab{
		crontab: gocron.NewScheduler(),
	}
}

// Every 方法用于在当前的定时任务中设置一个时间间隔。
// 这个方法接受一个无符号整数参数 interval，表示任务重复执行的时间间隔，单位为秒。
// 返回值是一个 *gocron.Job 对象，表示新创建的定时任务。
// 该方法主要通过调用 crontab 实例的 Every 方法来实现，本包提供了一层封装。
func (c *Crontab) Every(interval uint64) *gocron.Job {
	return c.crontab.Every(interval)
}

// Clear 清空Crontab任务列表
// 该方法通过调用crontab对象的Clear方法来清空所有已定义的定时任务
func (c *Crontab) Clear() {
	c.crontab.Clear()
}

// Start 启动一个Crontab任务调度器。
// 该方法通过接收来自Crontab实例的Start方法的信号来同步阻塞，
// 直到Crontab调度器成功启动或发生错误而停止。
// 这种设计模式确保了调用者可以知道调度器何时真正开始工作，
// 从而在调度器准备好之前避免进行任何操作。
func (c *Crontab) Start() {
	<-c.crontab.Start()
}
