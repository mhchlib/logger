package logger

import (
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
	"io"
)

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

func OutWrite() io.Writer {
	return logger.OutWrite()
}

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

func SetDebugLogLevel() {
	log.SetLevel(log.DebugLevel)
}
