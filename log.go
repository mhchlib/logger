package logger

import (
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"io"
)

// Error ...
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Info ...
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Debug ...
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Fatal ...
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// OutWrite ...
func OutWrite() io.Writer {
	return logger.OutWrite()
}

// PrintDataTable ...
func PrintDataTable(data [][]string, headers []string, desc string, opts ...func(table *tablewriter.Table)) {
	ioWriteInMermory := NewIoWriteMermory()
	//fist line is \n
	ioWriteInMermory.Write([]byte(desc + "\n"))
	table := tablewriter.NewWriter(ioWriteInMermory)
	table.SetHeader(headers)
	table.AppendBulk(data)
	for _, opt := range opts {
		opt(table)
	}
	table.Render()
	// mermory io write do not exist asynchronous write ， so just print
	logger.Info(ioWriteInMermory.PrintContent())
}

// SetDebugLogLevel ...
func SetDebugLogLevel() {
	log.SetLevel(log.DebugLevel)
}
