package logger

type IoWriteMermory struct {
	content []byte
}

func NewIoWriteMermory() *IoWriteMermory {
	return &IoWriteMermory{}
}

func (i *IoWriteMermory) Write(p []byte) (n int, err error) {
	i.content = append(i.content, p...)
	return len(p), nil
}

func (i *IoWriteMermory) PrintContent() string {
	return string(i.content)
}
