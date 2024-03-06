package cmclient

import "github.com/zeromicro/go-zero/core/logx"

// Logger define log interface
//type Logger interface {
//	Debugf(format string, args ...interface{})
//	Infof(format string, args ...interface{})
//	Warnf(format string, args ...interface{})
//	Errorf(format string, args ...interface{})
//
//	Debug(args ...interface{})
//	Info(args ...interface{})
//	Warn(args ...interface{})
//	Error(args ...interface{})
//}

//// cmcLogger 是你项目期望的日志接口
//type cmcLogger interface {
//	Debugf(format string, args ...interface{})
//	Infof(format string, args ...interface{})
//	Warnf(format string, args ...interface{})
//	Errorf(format string, args ...interface{})
//
//	Debug(args ...interface{})
//	Info(args ...interface{})
//	Warn(args ...interface{})
//	Error(args ...interface{})
//}

// LogAdapter 是 logx.Logger 到 utils.Logger 的适配器
type LogAdapter struct {
	log logx.Logger
}

// NewLogAdapter 创建 LogAdapter 实例
func NewLogAdapter(log logx.Logger) *LogAdapter {
	return &LogAdapter{log: log}
}

// Warnf 实现 utils.Logger 接口中的 Warnf 方法
func (la *LogAdapter) Warnf(format string, args ...interface{}) {
	la.log.Errorf(format, args...)
}

// Warn 实现 utils.Logger 接口中的 Warn 方法
func (la *LogAdapter) Warn(args ...interface{}) {
	la.log.Error(args...)
}

func (la *LogAdapter) Debugf(format string, args ...interface{}) {
	la.log.Debugf(format, args...)
}

func (la *LogAdapter) Infof(format string, args ...interface{}) {
	la.log.Infof(format, args...)
}

func (la *LogAdapter) Errorf(format string, args ...interface{}) {
	la.log.Errorf(format, args...)
}

func (la *LogAdapter) Debug(args ...interface{}) {
	la.log.Debug(args...)
}

func (la *LogAdapter) Info(args ...interface{}) {
	la.log.Info(args...)
}

func (la *LogAdapter) Error(args ...interface{}) {
	la.log.Error(args...)
}
