package options

import "time"

const (
    defaultType        = "node"
    defaultNonBlock    = true
    defaultPingTimeout = 1 * time.Second
)

type RedisConnectDemo struct {
    Host        string
    Type        string
    Pass        string
    Tls         bool
    NonBlock    bool
    PingTimeout time.Duration
}

type Option interface {
    apply(o *options)
}

type options struct {
    Host        string
    Type        string
    Pass        string
    Tls         bool
    NonBlock    bool
    PingTimeout time.Duration
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func WithHost(h string) Option {
    return optionFunc(func(o *options) { o.Host = h })
}

func WithNodeType(nodeType string) Option {
    return optionFunc(func(o *options) { o.Type = nodeType })
}

func WithPass(pass string) Option {
    return optionFunc(func(o *options) { o.Pass = pass })
}

func WithTls(tls bool) Option {
    return optionFunc(func(o *options) { o.Tls = tls })
}

func WithNonBlock(nonBlock bool) Option {
    return optionFunc(func(o *options) { o.NonBlock = nonBlock })
}

func WithPingTimeout(t time.Duration) Option {
    return optionFunc(func(o *options) { o.PingTimeout = t })
}

func NewRedisConnectDemo(opts ...Option) (*RedisConnectDemo, error) {
    _options := options{
        Type:        defaultType,
        NonBlock:    defaultNonBlock,
        PingTimeout: defaultPingTimeout,
    }

    for _, o := range opts {
        o.apply(&_options)
    }

    return &RedisConnectDemo{
        Host:        _options.Host,
        Type:        _options.Type,
        Pass:        _options.Pass,
        Tls:         _options.Tls,
        NonBlock:    _options.NonBlock,
        PingTimeout: _options.PingTimeout,
    }, nil
}
