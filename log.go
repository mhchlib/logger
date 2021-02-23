package logger

import log "github.com/sirupsen/logrus"

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func SetDebugLogLevel() {
	log.SetLevel(log.DebugLevel)
}
