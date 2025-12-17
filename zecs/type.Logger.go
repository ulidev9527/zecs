package zecs

import (
	"fmt"
	"time"
)

// ILogger 日志接口
type ILogger interface {
	Info(msg ...any)
	Warn(msg ...any)
	Error(msg ...any)
}

// Logger 控制台日志实现
type Logger struct {
	ILogger
}

func (c *Logger) log(level any, msg ...any) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] %s\n", timestamp, level, msg)
}

func (c *Logger) Info(msg ...any) {
	c.log("INFO", msg...)
}

func (c *Logger) Warn(msg ...any) {
	c.log("WARN", msg)
}

func (c *Logger) Error(msg ...any) {
	c.log("ERROR", msg)
}
