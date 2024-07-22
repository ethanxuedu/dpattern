package template

import "fmt"

type SmSer interface {
    manufacturer(mobile, code string) bool
    checkUserReceiveInOneMinute(mobile string) bool
    genCode() string
    saveRedis(mobile, code string)
}

// Sms 发送短信模板
// 定义了发送短信的基本流程操作
//
// 发送前检查用户 1 分钟内是否接收过短信，不变部分。
//
// 生成验证码，不变部分。
//
// 发远验证码到用户手机，这个抽象方法由不同子类实现，可变部分。
//
// 发送成功则保存到 redis 中，不变部分。
type Sms struct {
}

// manufacturer 由不同的厂商来实现发送短信到手机上
func (s *Sms) manufacturer(mobile, code string) bool {
    return true
}

// checkUserReceiveInOneMinute 检查1分钟内该手机号是否接收过验证码，1分钟内接收过就不能在发送验证码
func (s *Sms) checkUserReceiveInOneMinute(mobile string) bool {
    fmt.Println("该用户一分钟内没有接收过短信")
    return false
}

// genCode 生成6位验证码
func (s *Sms) genCode() string {
    return "123456"
}

// saveRedis 将手机号+验证码存进redis中，给登录接口做校验用
func (s *Sms) saveRedis(mobile, code string) {
    fmt.Println("将手机号加验证码存进 redis")
}

// HuaWei 具体子类 HuaWei 继承抽象类，实现抽象方法 manufacturer(String mobile, String code)，定义流程中的可变部分。
type HuaWei struct {
    Sms
}

// manufacturer 华为发送短信到手机上
func (h *HuaWei) manufacturer(mobile, code string) bool {
    fmt.Println("华为发送短信到手机上")
    return true
}

// Alibaba 具体子类 Alibaba 继承抽象类，实现抽象方法 manufacturer(String mobile, String code)，定义流程中的可变部分。
type Alibaba struct {
    Sms
}

// manufacturer 阿里发送短信到手机上
func (h *Alibaba) manufacturer(mobile, code string) bool {
    fmt.Println("阿里发送短信到手机上")
    return true
}

// Send 封装具体步骤,完成了基本流程组合与条件控制。
func Send(mobile string, s SmSer) {
    fmt.Println("检查用户一分钟内是否发送过短信，mobile:" + mobile)
    if s.checkUserReceiveInOneMinute(mobile) {
        fmt.Println("请等待1分钟后重试")
        return
    }
    code := s.genCode()
    if s.manufacturer(mobile, code) {
        fmt.Printf("厂商发送短信成功，mobile:%s,code:%s \n", mobile, code)
        s.saveRedis(mobile, code)
    }
}
