package main

import (
	l "learn-go-pro/log"
)

func main() {
	// baseLog := log.BaseLog{
	// 	LoggerLevel: log.DEBUG,
	// 	Output:      os.Stdout,
	// }
	// baseLog.Error("sss")
	// baseLog.Warning("sss")
	// baseLog.Info("sss")
	consoleLog := l.NewConsoleLog(l.DEBUG)

	consoleLog.Error("sss")
}
