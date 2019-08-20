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

// DoubleElement 列表结点
type DoubleElement struct {
	Data      Comparer
	pre, next *DoubleElement
}

// DoubleLink 列表
type DoubleLink struct {
	head   *DoubleElement
	lenght int
}

// NewDoubleList 新建双链表
func NewDoubleList() *DoubleLink {
	return &DoubleLink{
		head: &DoubleElement{},
	}
}

// Length 获取列表的长度
func (d *DoubleLink) Length() int {
	return d.lenght
}

// Insert 在列表的第几个位置插入元素
func (d *DoubleLink) Insert(i int, data Comparer) error {
	maxLen := d.Length() + 1
	if i <= 0 || i > maxLen {
		return ErrIndex
	}
	p := d.head

	for j := 1; j < i; j++ {
		p = p.next
	}
	p.next = &DoubleElement{
		Data: data,
		pre:  p,
		next: p.next,
	}
	d.lenght++
	return nil
}

// Delete 删除制定位置的元素
func (d *DoubleLink) Delete(i int) (data Comparer, err error) {
	if i <= 0 || i > d.Length() {
		return nil, ErrIndex
	}
	p := d.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	if p.next.next != nil {
		p.next.next.pre = p
	}
	data = p.next.Data
	p.next = p.next.next
	d.lenght--
	return data, nil
}

// PrintDoubleList 显示双向链表中的所有元素
func (d *DoubleLink) PrintDoubleList() string {
	lstr := "list:"
	for p := d.head.next; p != nil; {
		lstr += fmt.Sprint("\t", p.Data)
		p = p.next
	}
	fmt.Sprintln()
	return lstr
}

func main() {
	doubleList := NewDoubleList()
	doubleList.Insert(1, Age(1))
	doubleList.Insert(1, Age(2))
	doubleList.Insert(1, Age(3))
	doubleList.Insert(1, Age(4))
	fmt.Println(doubleList.PrintDoubleList())
}

type Age int

func (a Age) compare(data interface{}) bool {
	if a == data {
		return true
	}
	return false
}
