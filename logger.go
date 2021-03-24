package logger

import (
	"encoding/json"
	log "github.com/mhchlib/logger/internal"
	"io"
)

// Logger ...
type Logger interface {
	Info(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Debug(v ...interface{})
	Log(v ...interface{}) error
	OutWrite() io.Writer
}

// DefaultLogger ...
type DefaultLogger struct {
	opt      *Options
	internal bool
}

// NewLogger ...
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
		metaBytes, _ := json.Marshal(metaData)
		data = append(data, string(metaBytes))
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

// Info ...
func (d *DefaultLogger) Info(v ...interface{}) {
	extraData := d.getExtraData()
	v = append(extraData, v...)
	log.DoInfo(v...)
}

// Error ...
func (d *DefaultLogger) Error(v ...interface{}) {
	extraData := d.getExtraData()
	v = append(extraData, v...)
	log.DoError(v...)
}

// Fatal ...
func (d *DefaultLogger) Fatal(v ...interface{}) {
	extraData := d.getExtraData()
	v = append(extraData, v...)
	log.DoFatal(v...)
}

// Log ...
func (d *DefaultLogger) Log(v ...interface{}) error {
	d.Info(v)
	return nil
}

// Debug ...
func (d *DefaultLogger) Debug(v ...interface{}) {
	extraData := d.getExtraData()
	v = append(extraData, v...)
	log.DoDebug(v...)
}

// OutWrite ...
func (d *DefaultLogger) OutWrite() io.Writer {
	return d
}

// Write ...
func (d *DefaultLogger) Write(p []byte) (n int, err error) {
	log.DoPrint(string(p))
	return len(p), nil
}
