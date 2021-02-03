package list

import (
	"errors"
	"fmt"
)

const defaultCapacity = 10
const defaultSize = -1

//动态数组结构
//size	当前最后一个元素下标位置
//capacity	数组最大容量
//data	数组
type Array struct {
	size     int
	capacity int
	data     []interface{}
}

//创建一个新的动态数组
func NewArray(capacity int) *Array {
	data := make([]interface{}, capacity)
	return &Array{defaultSize, capacity, data}
}

//在不输入最大容量时创建
func NewArrayWithoutNoCap() *Array {
	data := make([]interface{}, defaultCapacity)
	return &Array{defaultSize, defaultCapacity, data}
}

//扩展数组
func (array *Array) grow(num int) (bool, []interface{}) {
	data := make([]interface{}, num)
	for i, elem := range array.data {
		data[i] = elem
	}
	return true, data
}

//检查数组容量是否足够
func (array *Array) check(num int) bool {
	if array.size+num+1 > array.capacity {
		ok, array1 := array.grow(array.size + num + 2)
		if ok == true {
			array.data = array1
			return true
		}
	}
	return false
}

//打印数组
func (array Array) Print() {
	for _, elem := range array.data {
		fmt.Print(elem)
		fmt.Print("\t")
	}
}

//向list末尾加入元素
func (array *Array) Add(obj ...interface{}) error {
	array.check(len(obj))
	for _, elem := range obj {
		array.size++
		array.data[array.size] = elem
	}
	return errors.New("no error")
}

//向指定位置加入元素
func (array *Array) Insert(location int, obj interface{}) error {
	array.check(1)
	for i := array.size; i >= location-1; i-- {
		array.data[i+1] = array.data[i]
	}
	array.data[location-1] = obj
	return errors.New("no error")
}

//向指定位置修改元素
func (array *Array) Set(location int, obj interface{}) error {
	if location == 0 || location > array.size+1 {
		return errors.New("下标超出")
	}
	array.data[location-1] = obj
	return errors.New("success")
}

//是否存在某元素
func (array *Array) Contain(obj interface{}) bool {
	for _, i := range array.data {
		if i == obj {
			return true
		}
	}
	return false
}

//是否为空
func (array *Array) IsEmpty() bool {
	if array.size == -1 {
		return true
	}
	return false
}

//查看某一位置上的元素
func (array *Array) Get(location int) (interface{}, error) {
	return array.data[location], errors.New("no error")
}

//判断是否相等
func (array *Array) Equals(array1 *Array) bool {
	for i, elem := range array.data {
		if array1.data[i] != elem {
			return false
		}
	}
	return true
}

//转换为Slice类型
func (array *Array) ToSlice() []interface{} {
	flag := make([]interface{}, array.capacity)
	for i, elem := range array.data {
		flag[i] = elem
	}
	return flag
}

//输出当前list的长度
func (array *Array) Size() int {
	return array.size + 1
}
