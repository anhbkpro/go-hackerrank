package logger_rate_limiter

// Logger Runtime: 136 ms, faster than 41.18% of Go online submissions for Logger Rate Limiter.
//Memory Usage: 10 MB, less than 31.62% of Go online submissions for Logger Rate Limiter.
type Logger struct {
	messages map[string]int
}

func Constructor() Logger {
	return Logger{
		messages: map[string]int{},
	}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if _, ok := l.messages[message]; !ok {
		l.messages[message] = timestamp
		return true
	}

	oldTimestamp := l.messages[message]
	if timestamp-oldTimestamp >= 10 {
		l.messages[message] = timestamp
		return true
	}

	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
