package factory

import "fmt"

type DemoAbstractFactory interface {
    PrintInfo()
}

type demoAbstractFactory struct {
    Name string
    Desc string
}

func (d *demoAbstractFactory) PrintInfo() {
    fmt.Printf("name is %s, desc is %s\n", d.Name, d.Desc)
}

// NewDemoAbstractFactory 和简单工厂模式的唯一区别，就是它返回的是接口而不是结构体。
// 通过返回接口，可以在你不公开内部实现的情况下，让调用者使用你提供的各种功能
// 如：demoAbstractFactory 是不公开的，但可以使用其提供的方法
func NewDemoAbstractFactory(name, desc string) DemoAbstractFactory {
    return &demoAbstractFactory{
        Name: name,
        Desc: desc,
    }
}
