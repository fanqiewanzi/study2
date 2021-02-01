package list

import (
	"errors"
	"fmt"
)

//单向链表
type LinkList struct {
	data interface{}
	next *LinkList
}

func (linkList *LinkList) Add(obj ...interface{}) error {
	//找最后一个元素
	p := linkList
	for p.next != nil {
		p = p.next
	}
	//在最后一个元素进行赋值，因为不知道有多少个值所以用循环
	for _, elem := range obj {
		q := new(LinkList)
		p.next = q
		q.data = elem
		p = q
	}
	return errors.New("no error")
}

func (linkList *LinkList) Insert(location int, obj interface{}) error {
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	if count != location {
		return errors.New("此位置超出限制")
	}
	//在相应位置元素进行赋值
	q := new(LinkList)
	q.next = p.next
	p.next = q
	q.data = obj
	return errors.New("no error")
}

func (linkList *LinkList) Set(location int, obj interface{}) error {
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	if count != location {
		return errors.New("此位置超出限制")
	}
	p.data = obj
	return errors.New("no error")
}

func (linkList *LinkList) Contain(obj interface{}) bool {
	//找相应位置元素
	p := linkList
	for p.next != nil && p.data != obj {
		p = p.next
	}
	if p.next == nil || p.data != obj {
		return false
	}
	return true
}

func (linkList *LinkList) IsEmpty() bool {
	if linkList.next == nil {
		return true
	} else {
		return false
	}
}

func (linkList *LinkList) Get(location int) (interface{}, error) {
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	if count != location {
		return nil, errors.New("此位置超出限制")
	}
	return p.data, errors.New("no error")
}

//判断是否相等
func (linkList *LinkList) Equals(linkList1 *LinkList) bool {
	//找相应位置元素
	p := linkList
	q := linkList1
	for p != nil && q != nil && p.data == q.data {
		p = p.next
		q = q.next
	}
	if p == nil && q == nil {
		return true
	}
	return false
}

//转换为Slice类型
func (linkList *LinkList) ToSlice() []interface{} {
	p := linkList
	var sli []interface{} = make([]interface{}, 50)
	i := 0
	for p != nil {
		sli[i] = p.data
		p = p.next
		i++
	}
	return sli
}

//输出当前list的长度
func (linkList *LinkList) Size() int {
	p := linkList
	count := 0
	for p != nil {
		count++
		p = p.next
	}
	return count
}

func (linkList *LinkList) Print() {
	p := linkList
	for p != nil {
		fmt.Print(p.data)
		p = p.next
	}
}
