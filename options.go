package logger

// Options ...
type Options struct {
	enableCodeData bool
	metaData       map[string]string
}

// LogOption ...
type LogOption func(*Options)

// EnableCodeData ...
func EnableCodeData(v bool) LogOption {
	return func(opts *Options) {
		opts.enableCodeData = v
	}
}

// MetaData ...
func MetaData(key string, value string) LogOption {
	return func(opts *Options) {
		metaData := opts.metaData
		if metaData == nil {
			metaData = make(map[string]string)
			opts.metaData = metaData
		}
		metaData[key] = value
	}
}
