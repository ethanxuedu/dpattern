package factory

import "fmt"

type DemoSimpleFactory struct {
    Name string
    Desc string
}

func (d *DemoSimpleFactory) PrintInfo() {
    fmt.Printf("name is %s, desc is %s\n", d.Name, d.Desc)
}

// NewDemoSimpleFactory 简单工厂模式，返回一个实例
// 和 p：=＆DemoFactory {}这种创建实例的方式相比，简单工厂模式可以确保我们创建的实例具有需要的参数，进而保证实例的方法可以按预期执行。
// 例如，通过 NewDemoFactory 创建 DemoFactory 实例时，可以确保实例的 name 和 desc 属性被设置。
// 注意：实际开发可以看情况返回指针与非指正类型
// 1. 非指针结构体时，Go 会创建这个结构体的副本并将副本返回给调用者。这意味着每次调用这个方法时，都会创建并复制整个结构体，这在某些情况下可能导致效率低下，尤其是结构体非常大且调用频繁的时候
// 2. 如果返回指针，Go 只需要复制指针（通常是一个内存地址），而不是整个结构体，这样可以减少内存和计算开销
func NewDemoSimpleFactory(name, desc string) *DemoSimpleFactory {
    return &DemoSimpleFactory{
        Name: name,
        Desc: desc,
    }
}
