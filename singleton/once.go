package singleton

import (
    "sync"
)

// 使用 once.Do 可以确保 ins 实例全局只被创建一次，once.Do 函数还可以确保当同时有多个创建动作时，只有一个创建动作在被执行。

type Singleton struct {
}

var ins *Singleton
var once sync.Once

func GetOnceInsOr() *Singleton {
    once.Do(func() {
        ins = &Singleton{}
    })
    return ins
}
