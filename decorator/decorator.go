package decorator

import (
    "fmt"
    "time"
)

type OSSer interface {
    Upload()
}

type TencentOss struct{}

func NewTencentOss() OSSer {
    return &TencentOss{}
}

func (t *TencentOss) Upload() {
    time.Sleep(time.Second)
    fmt.Println("上传文件到腾讯 oss 成功")
}

type HuaweiOss struct{}

func NewHuaweiOss() OSSer {
    return &TencentOss{}
}

func (h *HuaweiOss) Upload() {
    time.Sleep(time.Second * 2)
    fmt.Println("上传文件到华为 oss 成功")
}

type Decorator struct {
    oss        OSSer
    prometheus string
}

func NewDecorator(oss OSSer) *Decorator {
    return &Decorator{
        oss:        oss,
        prometheus: "统计了上传文件的执行时间",
    }
}

func (d *Decorator) Upload() {
    start := time.Now()
    defer func() {
        duration := time.Since(start).Milliseconds()
        fmt.Println(d.prometheus, ": ", duration)
    }()
    d.oss.Upload()
}
