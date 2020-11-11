package logger

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetReportCaller(true)
}

func Fatal(v ...interface{}) {
	extra := getExtraData()
	log.Fatal(extra, v)
}

func Error(v ...interface{}) {
	extra := getExtraData()
	log.Error(extra, v)
}

func Info(v ...interface{}) {
	extra := getExtraData()
	log.Info(extra, v)
}
