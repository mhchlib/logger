package logger

import (
	"github.com/micro/go-micro/v2/util/log"
)

func Fatal(v ...interface{}) {
	log.Fatal(v)
}

func Error(v ...interface{}) {
	log.Error(v)
}

func Println(v ...interface{}) {
	log.Info(v)
}
