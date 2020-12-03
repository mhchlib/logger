package logger

import "testing"

func TestInfo1(t *testing.T) {
	Info("111", "222", "333")
}

func TestInfo2(t *testing.T) {
	l := NewLogger(
		EnableCodeData(true),
		MetaData("key1", "value1"),
		MetaData("key2", "value2"),
		MetaData("key3", "value3"),
	)
	l.Info("111", "222", "333")
}
