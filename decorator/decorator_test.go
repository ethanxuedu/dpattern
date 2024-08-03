package decorator

import "testing"

func TestDecorator(t *testing.T) {
    tencentOss := NewTencentOss()
    d1 := NewDecorator(tencentOss)
    d1.Upload()

    huaweiOss := NewHuaweiOss()
    d2 := NewDecorator(huaweiOss)
    d2.Upload()
}
