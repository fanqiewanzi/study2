package singletonpattern

import "sync"

type Single struct {
	name string
}

//全局变量
var (
	single *Single
	once   sync.Once
)

func newSingle() *Single {
	//once.Do调用的函数只执行1次,所以这个是单例模式
	once.Do(func() {
		single = &Single{name: "BAZINGA"}
	})
	return single
}
