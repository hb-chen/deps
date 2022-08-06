package output

type Options struct {
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}
