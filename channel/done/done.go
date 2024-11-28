package main

import (
	"fmt"
)

// worker :用于从 Channel 中接收数据并处理
func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		//go func() { done <- true }()
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

// createWorker :这是一个工厂函数，用于创建每个 worker，并返回用于发送数据的 Channel。
func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done) //启动 Goroutine，运行 worker 函数，让它监听传入的 Channel。
	return w
}

// chanDemo :演示了如何使用多个 Channel 与多个 Goroutine 通信。
func chanDemo() {
	// 创建了一个 channels 数组，用于存储 10 个发送 Channel。
	var workers [10]worker
	for i := 0; i < 10; i++ {
		// 每个 Channel 对应一个 Worker Goroutine。
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	//for _, worker := range workers {
	//	<-worker.done
	//}

	//time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("channel as first-class citizen")
	chanDemo()
}
