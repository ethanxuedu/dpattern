package proxy

// 缓存代理可用于为开销大的操作结果提供缓存，以便多个客户端可以共享结果。
// 可以使用代理模式来实现缓存代理，以加快数据访问速度并减少对底层资源的频繁访问。

import "fmt"

type IGooder interface {
    FindGoodLimit100() string
}

type Good struct{}

func NewGood() IGooder {
    return &Good{}
}

func (g *Good) FindGoodLimit100() string {
    // select * from good order by id limit 100;
    return fmt.Sprintf("拿到 100 件商品数据返回")
}

type GoodCache struct {
    Good *Good
}

func NewGoodCache() IGooder {
    return &GoodCache{}
}

func (g *GoodCache) FindGoodLimit100() string {
    // 1. 从 cache 查询有没有缓存的商品
    // 2. cache 有商品，直接返回，缓存中没有商品，执行以下步骤
    // 3. select * from good order by id limit 100;
    res := g.Good.FindGoodLimit100()
    // 4. 从 db 中拿到 100 件商品放入 cache 中
    return fmt.Sprintf("db:%s", res)
}
