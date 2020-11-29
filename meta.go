package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func getCodeDesc(inside bool) string {
	skip := 0
	if inside {
		skip = 5
	} else {
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

func getTime() string {
	now := time.Now()
	format := now.Format("2006-01-02 15:04:05")
	return format
}

func getMetaData(inside bool) string {
	codeDesc := getCodeDesc(inside)
	time := getTime()
	return "[ " + time + " -> " + codeDesc + " ] "
}

func handleFileName(filename string) string {
	splits := strings.Split(filename, string(os.PathSeparator))
	len := len(splits)
	if len > 2 {
		filename = splits[len-2] + string(os.PathSeparator) + splits[len-1]
	}
	return filename
}

func handleFunName(funcName string) string {
	splits := strings.Split(funcName, string(os.PathSeparator))
	len := len(splits)
	if len > 1 {
		funcName = splits[len-1]
	}
	return funcName
}
