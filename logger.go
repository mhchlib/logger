package logger

import (
	log "github.com/sirupsen/logrus"
	"runtime"
)

func init() {
	//log.SetReportCaller(true)
}

func Fatal(v ...interface{}) {
	log.Fatal(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Println(v ...interface{}) {
	log.Info(v)

	pc, file, line, _ := runtime.Caller(1)
	f := runtime.FuncForPC(pc)
	log.Println(file, ":", line, ":", f.Name())
}
