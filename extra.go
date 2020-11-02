package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func getCodeDesc() string {
	pc, fileName, line, _ := runtime.Caller(3)
	f := runtime.FuncForPC(pc)
	funcName := f.Name()
	fileName = handleFileName(fileName)
	funcName = handleFunName(funcName)
	desc := fmt.Sprintf("%s:%d %s", fileName, line, funcName)
	return desc
}

func getExtraData() string {
	codeDesc := getCodeDesc()
	return "[ " + codeDesc + " ]  "
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