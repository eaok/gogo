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

// CircularElement 列表结点
type CircularElement struct {
	Data Comparer
	next *CircularElement
}

// CircularLink 列表
type CircularLink struct {
	head   *CircularElement
	lenght int
}

// NewCircularList 新建一个空列表
func NewCircularList() *CircularLink {
	head := &CircularElement{}
	head.next = head
	return &CircularLink{
		head: head,
	}
}

// Length 获取列表的长度
func (l *CircularLink) Length() int {
	return l.lenght
}

// Insert 在列表的第几个位置插入元素
func (l *CircularLink) Insert(i int, data Comparer) error {
	maxLen := l.Length() + 1
	if i <= 0 || i > maxLen {
		return ErrIndex
	}
	p := l.head
	for j := 1; j < i; j++ {
		p = p.next
	}
	p.next = &CircularElement{
		Data: data,
		next: p.next,
	}
	l.lenght++
	return nil
}

// Delete 删除制定位置的元素
func (l *CircularLink) Delete(i int) (data Comparer, err error) {
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

// PrintCircularList 显示循环链表中的元素
func (l *CircularLink) PrintCircularList() string {
	lstr := "list:"
	for p := l.head.next; p != l.head; {
		lstr += fmt.Sprint("\t", p.Data)
		p = p.next
	}
	fmt.Sprintln()
	return lstr
}

// GetElem 获取指定位置元素
func (l *CircularLink) GetElem(i int) (e *CircularElement, err error) {
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

// Reverse 反转列表
func (l *CircularLink) Reverse() error {
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
func (l *CircularLink) Union(sl *CircularLink) error {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = sl.head.next
	return nil
}

func main() {
	circularlist := NewCircularList()
	circularlist.Insert(1, Age(1))
	circularlist.Insert(1, Age(2))
	circularlist.Insert(1, Age(3))
	circularlist.Insert(1, Age(4))
	fmt.Println(circularlist.PrintCircularList())
	circularlist.Delete(1)
	fmt.Println(circularlist.PrintCircularList())

}

type Age int

func (a Age) compare(data interface{}) bool {
	if a == data {
		return true
	}
	return false
}
