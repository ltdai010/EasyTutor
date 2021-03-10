package logger

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"sync"
)

var (
	_logger *logs.BeeLogger
	once    sync.Once
)

func init() {
	once.Do(func() {
		onceInitLogger()
	})
}

func onceInitLogger() {
	fmt.Println("init logger")
	_logger = logs.NewLogger()
	_logger.SetLogger(logs.AdapterFile, `{"filename":"logs/EasyTutor.log","maxdays": 365,"perm":"0744"}`)
}

// Debug write log in debug level with format
func Debug(format string, v ...interface{}) {
	_logger.Debug(format, v...)
	fmt.Printf("[Debug] "+format+"\n", v...)
}

// Warn write log in warning level with format
func Warn(format string, v ...interface{}) {
	_logger.Warn(format, v...)
	fmt.Printf("[Warn] "+format+"\n", v...)
}

// Error write log in error level with format
func Error(format string, v ...interface{}) {
	_logger.Error(format, v...)
	fmt.Printf("[Error] "+format+"\n", v...)
}

// Info write log in info level with format
func Info(format string, v ...interface{}) {
	_logger.Info(format, v...)
	fmt.Printf("[Info] "+format+"\n", v...)
}
