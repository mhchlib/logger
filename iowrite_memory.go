package logger

// IoWriteMermory ...
type IoWriteMermory struct {
	content []byte
}

// NewIoWriteMermory ...
func NewIoWriteMermory() *IoWriteMermory {
	return &IoWriteMermory{}
}

// Write ...
func (i *IoWriteMermory) Write(p []byte) (n int, err error) {
	i.content = append(i.content, p...)
	return len(p), nil
}

// PrintContent ...
func (i *IoWriteMermory) PrintContent() string {
	return string(i.content)
}
