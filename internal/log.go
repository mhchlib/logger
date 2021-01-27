package internal

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func DoInfo(v ...interface{}) {
	log.Info(v)
}

func DoError(v ...interface{}) {
	log.Error(v)
}

func DoFatal(v ...interface{}) {
	log.Fatal(v)
}

func DoDebug(v ...interface{}) {
	log.Debug(v)
}
