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
	// once.Do 调用的函数只执行 1 次
	once.Do(func() {
		single = &Single{name: "Tom"}
	})
	return single
}
