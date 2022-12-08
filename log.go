package log

import "log"

const (
	// LOG_LEVEL_EMERGENCY  - System is unusable.
	// 系统不可用时
	LOG_LEVEL_EMERGENCY = iota
	
	// LOG_LEVEL_ALERT
	// Action must be taken immediately. Example: Entire website down, database unavailable, etc. This should trigger the SMS alerts and wake you up.
	// 必须立即采取行动时的报警。例如：整个网站无法访问，数据库不可用等。这应该触发短信警报并叫醒你。
	LOG_LEVEL_ALERT
	
	// LOG_LEVEL_CRITICAL
	// Critical conditions. Example: Application component unavailable, unexpected exception.
	// 关键条件时的报警。例如：应用程序组件不可用，意外异常。达到临界值时。
	LOG_LEVEL_CRITICAL
	
	// LOG_LEVEL_ERROR
	// Runtime errors that do not require immediate action but should typically be logged and monitored.
	LOG_LEVEL_ERROR
	
	// LOG_LEVEL_WARNING
	// Exceptional occurrences that are not errors. Example: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
	// 警告，出现异常，但不是错误，比如使用了过时的API，不推荐的API...
	LOG_LEVEL_WARNING
	
	// LOG_LEVEL_NOTICE
	// Normal but significant events.
	// 重要的通知
	LOG_LEVEL_NOTICE
	
	// LOG_LEVEL_INFO
	// Interesting events. Example: User logs in, SQL logs.
	// 信息
	LOG_LEVEL_INFO
	
	// LOG_LEVEL_DEBUG debug
	// Detailed debug information.
	// 调试信息
	LOG_LEVEL_DEBUG
)

type Logger struct {
	Logger   *log.Logger
	LogLevel []int
	Targets  []string
}
