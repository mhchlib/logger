package logger

type Logger interface {
	Info(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Log(keyvals ...interface{}) error
}

type DefaultLogger struct {
}

func (d DefaultLogger) Info(v ...interface{}) {
	doInfo(true, v)
}

func (d DefaultLogger) Error(v ...interface{}) {
	doError(true, v)
}

func (d DefaultLogger) Fatal(v ...interface{}) {
	doFatal(true, v)
}

func (d DefaultLogger) Log(v ...interface{}) error {
	d.Info(v)
	return nil
}
