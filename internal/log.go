package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
}

func DoInfo(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Info(vv...)
}

func DoError(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Error(vv...)
}

func DoFatal(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Fatal(vv...)
}

func DoDebug(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Debug(vv...)
}

func DoPrint(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	fmt.Print(vv...)
}
