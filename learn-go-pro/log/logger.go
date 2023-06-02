package log

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"runtime"
	"strings"
)

const (
	TRACE = iota

	DEBUG

	INFO

	WARNING

	ERROR

	FATAL
)

type LoggerLevel int

type Logger interface {
	Trace(format string, a ...interface{})

	Debug(format string, a ...interface{})

	Info(format string, a ...interface{})

	Warning(format string, a ...interface{})

	Error(format string, a ...interface{})

	Fatal(format string, a ...interface{})
}

type BaseLog struct {
	LoggerLevel
	Output *os.File
}

func (b BaseLog) Log(ll LoggerLevel, format string, a ...interface{}) {
	if ll < b.LoggerLevel {
		return
	}
	getInfo(2)
	_, err := fmt.Fprintf(b.Output, format, a...)
	if err != nil {
		panic(err)
	}
}

func (b BaseLog) Trace(format string, a ...interface{}) {
	b.Log(TRACE, format, a...)
}

func (b BaseLog) Debug(format string, a ...interface{}) {
	b.Log(DEBUG, format, a...)
}

func (b BaseLog) Info(format string, a ...interface{}) {
	b.Log(INFO, format, a...)
}

func (b BaseLog) Warning(format string, a ...interface{}) {
	b.Log(WARNING, format, a...)
}

func (b BaseLog) Error(format string, a ...interface{}) {
	b.Log(ERROR, format, a...)
}

func (b BaseLog) Fatal(format string, a ...interface{}) {
	b.Log(FATAL, format, a...)
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	regex := regexp.MustCompile(`[^/]+$`)
	println(regex.FindString(funcName))
	fileName = path.Base(file)
	funcName = strings.Split(funcName, ".")[1]
	return
}
