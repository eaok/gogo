package main

import "fmt"

const maxSize = 100

//SeqQueue 结构体定义
type SeqQueue struct {
	data  [maxSize]interface{}
	front int
	rear  int
}

//NewSeqQueue 新建空队列
func NewSeqQueue() *SeqQueue {
	return &SeqQueue{
		front: 0,
		rear:  0,
	}
}

//Length 队列长度
func (q *SeqQueue) Length() interface{} {
	return (q.rear - q.front + maxSize) % maxSize
}

//Enqueue 入队
func (q *SeqQueue) Enqueue(e interface{}) error {
	if (q.rear+1)%maxSize == q.front {
		return fmt.Errorf("quque is full")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % maxSize
	return nil
}

//Dequeue 出队
func (q *SeqQueue) Dequeue() (e interface{}, err error) {
	if q.rear == q.front {
		return e, fmt.Errorf("quque is empty")
	}
	e = q.data[q.front]
	q.front = (q.front + 1) % maxSize
	return e, nil
}

//PrintQueue 打印队列中的内容
func (q *SeqQueue) PrintQueue() {
	for i := q.front; i != q.rear; i = (i + 1) % maxSize {
		fmt.Print(q.data[i], " ")
	}
	fmt.Println()
}

func main() {
	queue := NewSeqQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.PrintQueue()
	fmt.Println(queue.Dequeue())
	queue.PrintQueue()
}
