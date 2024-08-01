package singleton

// 为了解决懒汉方式非并发安全的问题，需要对实例进行加锁，下面是带检查锁的一个实现：

import "sync"

type LLockSingleton struct{}

var (
    lLock *LLockSingleton
    mu    sync.Mutex
)

func GetLLockIns() *LLockSingleton {
    if lLock == nil {
        mu.Lock()
        if lLock == nil {
            lLock = &LLockSingleton{}
        }
        mu.Unlock()
    }
    return lLock
}

// 上述代码只有在创建时才会加锁，既提高了代码效率，又保证了并发安全。
