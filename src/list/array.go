package list

import "errors"

const defaultCapacity = 10
const defaultSize = -1

//动态数组结构
type Array struct {
	size     int
	capacity int
	data     []interface{}
}

//创建一个新的动态数组
func NewArray(size, capacity int, data []interface{}) *Array {
	return &Array{size, capacity, data}
}

//扩展数组
func (array *Array) Grow(num int) (bool, *Array) {
	return true, &Array{array.size, num, array.data}
}

//检查数组容量是否足够
func (array *Array) Check(num int) *Array {
	if array.size+num > array.capacity {
		ok, array := array.Grow((array.size + num) * 2)
		if ok == true {
			return array
		}
	}
	return array
}

//向list末尾加入元素
func (array *Array) Add(obj ...interface{}) error {
	array.Check(len(obj))
	for _, i := range obj {
		array.data[array.size] = i
		array.size++
	}
	return errors.New("no error")
}

//向指定位置加入元素
func (array *Array) Insert(location int, obj interface{}) error {
	array.Check(1)
	for i := array.size; i >= location; i-- {
		array.data[i+1] = array.data[i]
	}
	array.data[location] = obj
	return errors.New("no error")
}

//向指定位置修改元素
func (array *Array) Set(location int, obj interface{}) error {
	array.data[location] = obj
	return errors.New("no error")
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
	return array.size
}
