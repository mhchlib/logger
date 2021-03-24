package internal

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

// GetCodeDesc ...
func GetCodeDesc(internal bool) string {
	skip := 3
	if internal {
		skip = 4
	}
	pc, fileName, line, _ := runtime.Caller(skip)

	f := runtime.FuncForPC(pc)
	funcName := f.Name()
	fileName = handleFileName(fileName)
	funcName = handleFunName(funcName)
	desc := fmt.Sprintf("%s:%d %s", fileName, line, funcName)
	return desc
}

// GetTime ...
func GetTime() string {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	return format
}

func handleFileName(filename string) string {
	splits := strings.Split(filename, string(os.PathSeparator))
	length := len(splits)
	if length > 2 {
		filename = splits[length-2] + string(os.PathSeparator) + splits[length-1]
	}
	return filename
}

func handleFunName(funcName string) string {
	splits := strings.Split(funcName, string(os.PathSeparator))
	length := len(splits)
	if length > 1 {
		funcName = splits[length-1]
	}
	return funcName
}
