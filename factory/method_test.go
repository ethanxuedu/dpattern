package factory

import "testing"

// TestDemoMethodFactory1 没有使用工厂方法初始化一个默认年龄的工厂
func TestDemoMethodFactory1(t *testing.T) {
    defaultAge := int64(10)
    f1 := NewDemoMethodFactory("f1", defaultAge)
    f2 := NewDemoMethodFactory("f2", defaultAge)
    f3 := NewDemoMethodFactory("f3", defaultAge)

    t.Logf("%+v", f1)
    t.Logf("%+v", f2)
    t.Logf("%+v", f3)
}

// TestDemoMethodFactory2 有使用工厂方法初始化一个默认年龄的工厂
func TestDemoMethodFactory2(t *testing.T) {
    defaultAge := int64(10)
    newf := NewDefaultNameDemoMethodFactory(defaultAge)
    f1 := newf("f1")
    f2 := newf("f2")
    f3 := newf("f3")

    t.Logf("%+v", f1)
    t.Logf("%+v", f2)
    t.Logf("%+v", f3)
}

func TestNewDefaultMarkDemoMethodFactory(t *testing.T) {
    defaultMark := "mark:"
    newFunc := NewDefaultMarkDemoMethodFactory(defaultMark)
    f1 := newFunc("f1", 1)
    f2 := newFunc("f2", 2)
    f3 := newFunc("f3", 3)

    t.Logf("%+v", f1)
    t.Logf("%+v", f2)
    t.Logf("%+v", f3)
}
