package logger

var logger Logger

func init() {
	logger = NewLogger(
		EnableCodeData(true),
	)
	logger.(*DefaultLogger).enableInternalLogger()
}
