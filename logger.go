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
	if d.opt.enableCodeData {
		codeDesc = log.GetCodeDesc(d.internal)
	}
	metaData := d.opt.metaData
	time := log.GetTime()
	return []interface{}{time, metaData, codeDesc}
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
