package logger

import (
	log "github.com/mhchlib/logger/internal"
)

type Logger interface {
	Info(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Log(v ...interface{}) error
}

type DefaultLogger struct {
	opt      *Options
	internal bool
}

func NewLogger(opts ...LogOption) Logger {
	return newLogger(opts)
}

func newLogger(options []LogOption) Logger {
	opts := &Options{}
	for _, o := range options {
		o(opts)
	}
	logger := &DefaultLogger{
		opt: opts,
	}
	return logger
}

func (d *DefaultLogger) getExtraData() []interface{} {
	var codeDesc = ""
	metaData := d.opt.metaData
	time := log.GetTime()
	data := []interface{}{time}
	if len(metaData) != 0 {
		data = append(data, metaData)
	}
	if d.opt.enableCodeData {
		codeDesc = log.GetCodeDesc(d.internal)
		data = append(data, codeDesc)
	}
	return data
}

func filter(data []interface{}) []interface{} {
	newData := []interface{}{}
	for _, v := range data {
		if v != nil {
			newData = append(newData, v)
		}
	}
	if len(newData) == 0 {
		return nil
	}
	return newData
}

func (d *DefaultLogger) enableInternalLogger() {
	d.internal = true
}

func (d *DefaultLogger) Info(v ...interface{}) {
	extraData := d.getExtraData()
	log.DoInfo(extraData, v)
}

func (d *DefaultLogger) Error(v ...interface{}) {
	extraData := d.getExtraData()
	log.DoError(extraData, v)
}

func (d *DefaultLogger) Fatal(v ...interface{}) {
	extraData := d.getExtraData()
	log.DoFatal(extraData, v)
}

func (d *DefaultLogger) Log(v ...interface{}) error {
	d.Info(v)
	return nil
}
