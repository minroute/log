package log

type TargetInstance struct {
	// Target is the target to which the log messages are written.
	Instance *Target
}

type Target interface {
	
	// Init initailize the target
	// 初始化输出目标
	Init(v ...interface{}) (*TargetInstance, error)
	
	// SetLevel 设置该输出的日志级别，具体日志级别请参考：LEVEL_LOG_*
	// 只有符合该级别的日志才会被输出
	// @see LEVEL_LOG_* for more information  on log.go file.
	SetLevel(level []int) (*TargetInstance, error)
	
	// SetGroup 设置日志分组,支持*通配符，支持多个分组
	// 当日志消息指定分组的时候，进行模式匹配，匹配成功则输出，否则不输出
	SetGroup(group []string) (*TargetInstance, error)
	
	// SetFlushInterval 设置日志刷新间隔,单位秒
	// Info: Message flushing also occurs when the application ends, which ensures log targets can receive complete log messages.
	SetFlushInterval(interval int) (*TargetInstance, error)
	
	// GetMessages - Get a list of log messages ([[\Yii\Log\Message]] instances).
	GetMessages() []*string
	
	// GetFormattedMessages - Get a list of log messages formatted as strings.
	GetFormattedMessages() []string
	
	// FormatMessages - Get all log messages formatted as a string.
	FormatMessages(messages []*string) string
	
	// GetCommonContext - Get an array with common context data in the key => value format.
	GetCommonContext() map[string]interface{}
}
