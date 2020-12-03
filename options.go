package logger

type Options struct {
	enableCodeData bool
	metaData       map[string]string
}

type LogOption func(*Options)

func EnableCodeData(v bool) LogOption {
	return func(opts *Options) {
		opts.enableCodeData = v
	}
}

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
