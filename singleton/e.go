package singleton

// 饿汉模式：实例是在包被导入时初始化的，所以如果初始化耗时，会导致程序加载时间比较长。

type ESingleton struct{}

var e *ESingleton = &ESingleton{}

func GetEInsOr() *ESingleton {
    return e
}
