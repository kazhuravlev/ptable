package ptable

type options struct {
	IncludeFields []string
}

type optionFn func(opt *options)

func WithIncludeFields(fields ...string) optionFn {
	return func(opt *options) {
		opt.IncludeFields = fields
	}
}
