package factory

type DemoMethodFactory struct {
    Mark string
    Age  int64
    Name string
}

func NewDemoMethodFactory(name string, age int64) *DemoMethodFactory {
    return &DemoMethodFactory{
        Name: name,
        Age:  age,
    }
}

func NewDefaultNameDemoMethodFactory(age int64) func(name string) DemoMethodFactory {
    return func(name string) DemoMethodFactory {
        return DemoMethodFactory{Age: age, Name: name}
    }
}

// NewDefaultMarkDemoMethodFactory 在简单工厂模式中，依赖于唯一的工厂对象，如果我们需要实例化一个产品，就要向工厂中传入一个参数，
// 获取对应的对象；如果要增加一种产品，就要在工厂中修改创建产品的函数。这会导致耦合性过高，这时我们就可以使用工厂方法模式。
// 在工厂方法模式中，依赖工厂函数，我们可以通过实现工厂函数来创建多种工厂，将对象创建从由一个对象负责所有具体类的实例化，
// 变成由一群子类来负责对具体类的实例化，从而将过程解耦。
// 比如: 突然要往结构体增加一个默认 Mark string 字段，每个结构体的都需要这个默认的字段，可以往结构体增加这个字段，再写以下工厂方法
func NewDefaultMarkDemoMethodFactory(mark string) func(name string, age int64) DemoMethodFactory {
    return func(name string, age int64) DemoMethodFactory {
        return DemoMethodFactory{Mark: mark, Age: age, Name: name}
    }
}
