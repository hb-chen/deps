package template

type Options struct {
	TplFile string
	OutFile string
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithTplFile(path string) Option {
	return func(o *Options) {
		o.TplFile = path
	}
}

func WithOutFile(path string) Option {
	return func(o *Options) {
		o.OutFile = path
	}
}
