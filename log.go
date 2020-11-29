package logger

import (
	log "github.com/sirupsen/logrus"
)

func Fatal(v ...interface{}) {
	doFatal(false, v)
}

func Error(v ...interface{}) {
	doError(false, v)
}

func Info(v ...interface{}) {
	doInfo(false, v)
}

func doInfo(inside bool, v ...interface{}) {
	extra := getMetaData(inside)
	log.Info(extra, v)
}

func doError(inside bool, v ...interface{}) {
	extra := getMetaData(inside)
	log.Error(extra, v)
}

func doFatal(inside bool, v ...interface{}) {
	extra := getMetaData(inside)
	log.Fatal(extra, v)
}
