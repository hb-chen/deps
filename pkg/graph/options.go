package graph

type Options struct {
	ProjectPath string
}

type Option func(*Options)

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

func WithProjectPath(path string) Option {
	return func(o *Options) {
		o.ProjectPath = path
	}
}
