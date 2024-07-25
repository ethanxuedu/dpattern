package factory

import "testing"

// PrintFactoryInfo 打印工厂信息
func PrintFactoryInfo(f DemoAbstractFactory) {
    f.PrintInfo()
}

func TestPrintFactoryInfo(t *testing.T) {
    f := NewDemoAbstractFactory("f1", "f1 is abstract factory pattern test demo")

    // 假设想执行 PrintFactoryInfo 来测试
    PrintFactoryInfo(f)
}
