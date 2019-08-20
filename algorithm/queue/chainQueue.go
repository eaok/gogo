package main

import (
	"fmt"
)

// Element is data
type Element int

// Queue is Queue Data
type Queue struct {
	e    Element
	next *Queue
}

// QueueLink 是队列
type QueueLink struct {
	front, rear *Queue
	length      int
}

//New 新建队列
func NewQueueLink() *QueueLink {
	q := Queue{e: Element(0)}
	return &QueueLink{front: &q, rear: &q, length: 0}
}

//Clear 清空队列
func (q *QueueLink) Clear() bool {
	return true
}

//Destroy 销毁队列
func (q *QueueLink) Destroy() bool {
	return true
}

//Empty 队列是否为空
func (q *QueueLink) Empty() bool {
	if q.front == q.rear {
		return true
	}
	return false
}

//Head 获取队头元素
func (q *QueueLink) Head() (e Element, err error) {
	if q.front == q.rear {
		err = fmt.Errorf("queue is empty")
		return
	}
	e = q.front.next.e
	return
}

//Enqueue : 进队列
func (q *QueueLink) Enqueue(e Element) (err error) {
	node := Queue{e: e}
	q.rear.next = &node
	q.rear = q.rear.next
	q.length++
	return
}

//Dequeue : 出队列
func (q *QueueLink) Dequeue() (e Element, err error) {
	if q.front == q.rear {
		err = fmt.Errorf("queue is empty")
		return
	}
	e = q.front.next.e
	q.front = q.front.next.next
	q.length--
	if q.length == 0 {
		q.rear = q.front
	}
	return
}

//Length 获取队列长度
func (q *QueueLink) Length() int {
	return q.length
}

//PrintQueue 打印队列中的内容
func (q *QueueLink) PrintQueue() {
	for i := q.front; i != q.rear; i = i.next {
		fmt.Print(i.next.e, " ")
	}
	fmt.Println()
}

func main() {
	queue := NewQueueLink()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.PrintQueue()
}