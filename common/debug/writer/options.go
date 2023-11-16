/*
 * @Author: 尤_Ta
 * @Date:  23:32
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:32
 */

package writer

// Options 可配置参数
type Options struct {
	path   string
	suffix string //文件扩展名
	cap    uint
}

func setDefault() Options {
	return Options{
		path:   "/tmp/go-admin",
		suffix: "log",
	}
}

// Option set options
type Option func(*Options)

// WithPath set path
func WithPath(s string) Option {
	return func(o *Options) {
		o.path = s
	}
}

// WithSuffix set suffix
func WithSuffix(s string) Option {
	return func(o *Options) {
		o.suffix = s
	}
}

// WithCap set cap
func WithCap(n uint) Option {
	return func(o *Options) {
		o.cap = n
	}
}
