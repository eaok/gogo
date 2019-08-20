package main

import (
	"errors"
	"fmt"
)

var (
	//ErrIndex 超出列表索引
	ErrIndex = errors.New("out of list index")
	//ErrNotFound 没有找到该元素
	ErrNotFound = errors.New("not found this element")
)

// Comparer 数据接口
type Comparer interface {
	compare(data interface{}) bool
}

// SingleElement 列表结点
type SingleElement struct {
	Data Comparer
	next *SingleElement
}

// SingleLink 列表
type SingleLink struct {
	head   *SingleElement
	lenght int
}

// NewSingleList 新建一个空列表
func NewSingleList() *SingleLink {
	return &SingleLink{
		head: &SingleElement{},
	}
}

// Length 获取列表的长度
func (l *SingleLink) Length() int {
	return l.lenght
}

// Insert 在列表的第几个位置插入元素
func (l *SingleLink) Insert(i int, data Comparer) error {
	maxLen := l.Length() + 1
	if i <= 0 || i > maxLen {
		return ErrIndex
	}
	p := l.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	p.next = &SingleElement{
		Data: data,
		next: p.next,
	}
	l.lenght++
	return nil
}

// Delete 删除制定位置的元素
func (l *SingleLink) Delete(i int) (data Comparer, err error) {
	if i <= 0 || i > l.Length() {
		return nil, ErrIndex
	}
	p := l.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	data = p.next.Data
	p.next = p.next.next
	l.lenght--
	return data, nil
}

// PrintSingleList 显示链表中所有元素
func (l *SingleLink) PrintSingleList() string {
	lstr := "list:"
	for p := l.head.next; p != nil; {
		lstr += fmt.Sprint("\t", p.Data)
		p = p.next
	}
	fmt.Sprintln()
	return lstr
}

// GetElem 获取指定位置元素
func (l *SingleLink) GetElem(i int) (e *SingleElement, err error) {
	if i <= 0 || i > l.Length() {
		return nil, ErrIndex
	}
	p := l.head
	for j := 1; j <= i; j++ {
		p = p.next
	}
	e = p
	return
}

// LocateElem 获取指定值的索引
func (l *SingleLink) LocateElem(data Comparer) (i int, err error) {
	p := l.head.next
	for p != nil {
		i++
		if p.Data.compare(data) {
			return i, nil
		}
		p = p.next
	}
	i = 0
	err = ErrNotFound
	return
}

// Reverse 反转列表
func (l *SingleLink) Reverse() error {
	if l.Length() == 0 {
		return nil
	}
	p := l.head.next
	pre := p.next
	p.next = nil

	for pre != nil {
		t := pre.next
		pre.next = p
		p, pre = pre, t
	}

	l.head.next = p
	return nil
}

// Union 两个链表合并
func (l *SingleLink) Union(sl *SingleLink) error {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = sl.head.next
	l.lenght += sl.lenght
	return nil
}

func main() {
	singleList := NewSingleList()
	singleList.Insert(1, Age(1))
	singleList.Insert(1, Age(2))
	singleList.Insert(1, Age(3))
	singleList.Insert(1, Age(3))
	fmt.Println(singleList.GetElem(1))
	fmt.Println(singleList.GetElem(2))
	fmt.Println(singleList.PrintSingleList())
	singleList.Delete(1)
	fmt.Println(singleList.PrintSingleList())
	singleList.Reverse()
	fmt.Println(singleList.PrintSingleList())
}

type Age int

func (a Age) compare(data interface{}) bool {
	if a == data {
		return true
	}
	return false
}
