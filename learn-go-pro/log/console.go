package log

import (
	"os"
)

type ConsoleLog struct {
	BaseLog
}

func NewConsoleLog(level LoggerLevel) ConsoleLog {
	return ConsoleLog{
		BaseLog{
			LoggerLevel: level,
			Output:      os.Stdout,
		},
	}
}
