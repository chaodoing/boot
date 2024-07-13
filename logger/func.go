package logger

import (
	`io`
	`log`
	`os`
	`path`
	`time`
	
	`github.com/gookit/goutil/fsutil`
	`github.com/kataras/golog`
	`github.com/lestrrat-go/strftime`
	`gopkg.in/natefinch/lumberjack.v2`
	`gorm.io/gorm/logger`
)

func (l *Logger) Writer() (write io.Writer, err error) {
	// 获取日志文件所在目录
	var dir = path.Dir(os.ExpandEnv(l.File))
	if !fsutil.PathExists(dir) {
		if err = fsutil.Mkdir(dir, 0755); err != nil {
			return
		}
	}
	p, err := strftime.New(os.ExpandEnv(l.File))
	if err != nil {
		return
	}
	write = &lumberjack.Logger{
		Filename:   p.FormatString(time.Now()), // 设置日志文件名，使用时间戳区分
		MaxSize:    4,                          // 最大尺寸，单位MB
		MaxAge:     1,                          // 最长保留时间，单位天
		MaxBackups: 31,                         // 最多备份文件数量
		LocalTime:  true,                       // 使用本地时间
		Compress:   true,                       // 是否压缩备份文件
	}
	if l.Stdout {
		write = io.MultiWriter(write, os.Stdout)
	}
	return
}

func (l *Logger) Log() (Log *log.Logger, err error) {
	var write io.Writer
	write, err = l.Writer()
	if err != nil {
		return nil, err
	}
	Log = log.New(write, "", log.LstdFlags|log.Ldate|log.Ltime)
	return
}

func (l *Logger) IrisLevel() string {
	return golog.Level(l.Level).String()
}

func (l *Logger) GormLevel() logger.LogLevel {
	return logger.LogLevel(l.Level)
}
