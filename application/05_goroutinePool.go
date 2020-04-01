package main

import (
	"fmt"
	"time"
)

// Task 有关Task任务相关定义及操作
type Task struct {
	f func() error //一个无参的函数类型
}

// NewTask 创建一个Task
func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}

	return &t
}

// Execute 调用任务所绑定的函数
func (t *Task) Execute() {
	t.f()
}

// Pool 定义协程池类型
type Pool struct {
	workerNum   int        //协程池最大worker数量,限定Goroutine的个数
	JobsChannel chan *Task //协程池内部的任务就绪队列
}

// NewPool 创建一个协程池
func NewPool(cap int) *Pool {
	p := Pool{
		workerNum:   cap,
		JobsChannel: make(chan *Task),
	}

	return &p
}

// worker 协程池创建一个worker并且开始工作
func (p *Pool) worker(workId int) {
	//worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel {
		//如果拿到任务,则执行task任务
		task.Execute()
		fmt.Println("worker ID ", workId, " 执行完毕任务")
	}
}

// Run 让协程池Pool开始工作
func (p *Pool) Run() {
	// 开启固定数量的Worker,每一个Worker用一个Goroutine承载
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
}

func main() {
	//创建一个Task
	t0 := NewTask(func() error {
		fmt.Println(0)
		return nil
	})

	t1 := NewTask(func() error {
		fmt.Println(1)
		time.Sleep(time.Second)
		return nil
	})

	t2 := NewTask(func() error {
		fmt.Println(2)
		time.Sleep(time.Second * 2)
		return nil
	})

	//创建一个协程池,最大开启3个协程worker
	p := NewPool(3)
	//启动协程池p
	p.Run()

	//添加task任务
	p.JobsChannel <- t0

	go func() {
		for {
			p.JobsChannel <- t1
		}
	}()
	go func() {
		for {
			p.JobsChannel <- t2
		}
	}()

	defer func() {
		close(p.JobsChannel)
	}()

	time.Sleep(time.Second * 5)
}
