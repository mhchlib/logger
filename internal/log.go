package internal

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func init() {
}

// DoInfo ...
func DoInfo(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Info(vv...)
}

// DoError ...
func DoError(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Error(vv...)
}

// DoFatal ...
func DoFatal(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Fatal(vv...)
}

// DoDebug ...
func DoDebug(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	log.Debug(vv...)
}

// DoPrint ...
func DoPrint(v ...interface{}) {
	var vv []interface{}
	for _, i := range v {
		vv = append(vv, i, " ")
	}
	fmt.Print(vv...)
}
