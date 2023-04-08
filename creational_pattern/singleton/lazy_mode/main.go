package main

import (
	"fmt"
	"sync"
	"time"
)

type singletonNotSafe struct {}

var instance *singletonNotSafe
var num int

func GetInstanceNotSafe() *singletonNotSafe {
	//只有首次GetInstance()方法被调用，才会生成这个单例的实例
	if instance == nil {
		// 用这个来模拟创建对象的耗时，否则因为程序执行太快，模拟不出争抢的现象
		time.Sleep(time.Nanosecond)
		num += 10
		instance = new(singletonNotSafe)
		return instance
	}

	//接下来的GetInstance直接返回已经申请的实例即可
	return instance
}

func (s *singletonNotSafe) SomeThing() {
	fmt.Println("单例对象的某方法")
}


type lazySingleton struct {}

func (r lazySingleton) SomeMethod()  {
	fmt.Println("单例的某些方法")
}


var LazySingleton *lazySingleton
var once sync.Once

// GetLazyInstance 在具体调用此方法时进行初始化
func GetLazyInstance() *lazySingleton {
	// Do 保证其中的代码只会执行一次。底层是使用的 atomic 进行控制的
	once.Do(func() {
		// 用这个来模拟创建对象的耗时，否则因为程序执行太快，模拟不出争抢的现象
		time.Sleep(time.Nanosecond)
		num += 1
		LazySingleton = &lazySingleton{}
	})

	return LazySingleton
}


func main()  {

	for i := 0; i < 20; i++ {
		go GetInstanceNotSafe()

	}
	time.Sleep(time.Second)
	fmt.Println(num)

}