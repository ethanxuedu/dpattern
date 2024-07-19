package options

import (
    "testing"
    "time"
)

func TestNewRedisConnectDemo(t *testing.T) {
    host, _type, pass := "127.0.0.1:6379", "node", "123456"
    tls, nonBlock := true, true
    pingTimeout := 1 * time.Second

    connect, err := NewRedisConnectDemo(
        WithHost(host),
        WithNodeType(_type),
        WithPass(pass),
        WithTls(tls),
        WithNonBlock(nonBlock),
        WithPingTimeout(pingTimeout),
    )

    if err != nil {
        t.Error(err)
        return
    }

    t.Logf("%+v", connect)
}

func TestNewRedisConnectDemo2(t *testing.T) {
    host, _type, pass := "127.0.0.1:6379", "node", "123456"
    tls, nonBlock := true, true
    pingTimeout := 1 * time.Second

    opts := []Option{
        WithHost(host),
        WithNodeType(_type),
        WithPass(pass),
        WithTls(tls),
        WithNonBlock(nonBlock),
        WithPingTimeout(pingTimeout),
    }
    connect, err := NewRedisConnectDemo(opts...)

    if err != nil {
        t.Error(err)
        return
    }

    t.Logf("%+v", connect)
}
