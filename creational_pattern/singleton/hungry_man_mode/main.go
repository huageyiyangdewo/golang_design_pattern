package main

// singleton 是单例模式
type singleton struct {}


var instance *singleton

// init 函数在 import singleton 包的时候就会执行其中的代码
func init()  {
	instance = &singleton{}
}

func GetInstance() *singleton {
	return instance
}

