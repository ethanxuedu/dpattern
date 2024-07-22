package template

import "testing"

func TestSmSer(t *testing.T) {
    phone := "18918998989"
    huawei := &HuaWei{}
    Send(phone, huawei)

    t.Log("-------------------")

    alibaba := &Alibaba{}
    Send(phone, alibaba)
}
