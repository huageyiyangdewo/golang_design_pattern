package main

import "fmt"

//=========实现层=============

// PayConfig 支付配置类，这里是做一个假设，实际上的支付配置比这多很多
type PayConfig struct {
	Name string
	AppId string
	NotifyUrl string
	SingType string
}

// GetConf 获取具体的配置信息
func (p *PayConfig)  GetConf()  {
	switch p.Name {
	case "ali":
		p.AppId = "xxx"
		p.SingType = "RSA"
		p.NotifyUrl = "http://xxx"
		fmt.Println("支付宝支付")
	case "wechat":
		p.AppId = "yyy"
		p.NotifyUrl = "http://yyyy"
		fmt.Println("微信支付")
	}
}

func NewPayConf(name string) (p *PayConfig) {
	switch name {
	case "ali":
		p = &PayConfig{Name: name}
	case "wechat":
		p = &PayConfig{Name: name}
	}

	p.GetConf()
	return p
}

//=========业务逻辑层=============

func main()  {
	p := new(PayConfig)

	p = NewPayConf("ali")

	p = NewPayConf("wechat")

	fmt.Println(p)
}