package main

import "fmt"

//=========抽象层=============

type PayConfer interface {
	GetConf()
}

//=========实现层=============

// CommonPayConfig 支付通用配置类，这里是做一个假设，实际上的支付配置比这多很多
type CommonPayConfig struct {
	Name string
	AppId string
	NotifyUrl string
}

type AliPayConfig struct {
	CommonPayConfig
	SingType string
}

func (a *AliPayConfig) GetConf() {
	a.Name = "ali"
	a.AppId = "xxx"
	a.SingType = "RSA"
	a.NotifyUrl = "http://xxx"
	fmt.Println("支付宝支付")
}

type WechatPayConf struct {
	CommonPayConfig
}

func (w *WechatPayConf) GetConf() {
	// 假设是从 ini 配置文件中获取
	w.Name = "wechat"
	w.AppId = "yyy"
	w.NotifyUrl = "http://yyyy"
	fmt.Println("微信支付")
}

type Factory struct {}

func (f *Factory) CreatePayConf(name string) (p PayConfer) {
	switch name {
	case "ali":
		// 假设是从 yaml 配置文件中获取
		p = &AliPayConfig{}
	case "wechat":
		// 假设是从 ini 配置文件中获取
		p = &WechatPayConf{}
	}
	p.GetConf()
	return p
}


//=========业务逻辑层=============
func main()  {

	p := new(Factory)
	ali := p.CreatePayConf("ali")
	fmt.Printf("ali -- %#v\n", ali)

	wechat := p.CreatePayConf("wechat")
	fmt.Printf("%#v\n", wechat)

}


