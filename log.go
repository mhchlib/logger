package logger

func Error(v ...interface{}) {
	logger.Error(v)
}

func Info(v ...interface{}) {
	logger.Info(v)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v)
}
