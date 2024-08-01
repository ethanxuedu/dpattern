package singleton

// 懒汉方式是开源项目中使用最多的，但它的缺点是非并发安全，在实际使用时需要加锁。以下是懒汉方式不加锁的一个实现：

type LSingleton struct{}

var l *LSingleton

func GetInsOr() *LSingleton {
    if l == nil {
        l = &LSingleton{}
    }
    return l
}

// 可以看到，在创建 l 时，如果 l==nil，就会再创建一个 l 实例，并发时就会有多个实例。
