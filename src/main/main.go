package main

import (
	"calculate"
	"fmt"
	"list"
	"mix"
)

func main() {
	mix.Count("str324ikasjdh 1234+)sd<>")
	mix.Fibonacci()
	fmt.Println("\n链表测试")
	linkTest()
	fmt.Println("\n数组测试")
	arraylistTest()
	fmt.Println("\n计算器测试")
	calculateTest()
	return
}

func linkTest() {
	list1 := list.NewLinkList()
	list2 := list.NewLinkList()
	list1.Add(1, 2, 3, 4)
	list2.Add(1, 2, 3, 4)
	list1.Print()
	list1.Insert(2, 5)
	list1.Print()
	fmt.Println(list1.Contain(5))
	list1.Set(3, 1)
	list1.Print()
	fmt.Println(list1.IsEmpty())
	fmt.Println(list1.Get(4))
	fmt.Println(list1.Equals(list2))
	fmt.Println(list1.ToSlice())
	fmt.Println(list1.Size())
}
func calculateTest() {
	calculate.Calculate("(2+3)/6+1")
}

func arraylistTest() {
	list1 := list.NewArray(10)
	list2 := list.NewArray(10)
	list1.Add(1, 2, 3, 4)
	list1.Print()
	list1.Insert(3, 1)
	list1.Set(2, 4)
	fmt.Println(list1.Equals(list2))
	fmt.Println(list1.Contain(2))
	fmt.Println(list1.IsEmpty())
	fmt.Println(list1.Size())
	fmt.Println(list1.ToSlice())
	fmt.Println(list1.Get(2))
	list1.Add(243, 43, 34, 3, 2, 2, 12, 4)
	list1.Print()
	fmt.Println()
	fmt.Print(list1.Size())
}
