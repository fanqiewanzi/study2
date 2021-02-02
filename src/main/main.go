package main

import (
	"calculate"
	"fmt"
	"list"
	"mix"
)

func main() {
	str := "sdkja3al123++  sadf"
	mix.Count(str)
	mix.Fibonacci()
	linkTest()
	calculate.Calculate("32+")

	return
}

func linkTest() {
	list1 := new(list.LinkList)
	list2 := new(list.LinkList)
	list1.Add(1, 2, 3, 4)
	list2.Add(1, 2, 3, 4)
	fmt.Println()
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
	calculate.Calculate("23/")
}
