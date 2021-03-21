package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
}

func DoInfo(v ...interface{}) {
	log.Info(v...)
}

func DoError(v ...interface{}) {
	log.Error(v...)
}

func DoFatal(v ...interface{}) {
	log.Fatal(v...)
}

func DoDebug(v ...interface{}) {
	log.Debug(v...)
}

func DoPrint(v ...interface{}) {
	fmt.Print(v...)
}
