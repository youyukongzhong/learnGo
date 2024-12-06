package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generator 函数创建一个生成整数的生产者。
// 它每隔一定时间（随机时间）生成一个整数，并通过 channel 返回。
// 该函数运行在一个 goroutine 中，持续生成数据并发送到 channel 中。
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		// 无限循环，持续生成数据
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			// 发送数据到 channel
			out <- i
			i++ // 每次生成的数据递增
		}
	}()
	return out // 返回生成数据的 channel
}

// worker 函数代表消费者，每个 worker 从 channel 中接收数据并处理。
// 每个 worker 在接收到数据后会休眠一秒，然后打印处理的结果。
func worker(id int, c chan int) {
	//遍历 Channel，直到 Channel 被关闭。
	//每次接收一个数据 n，输出到控制台。
	//range 的作用：优雅地处理 Channel 的关闭，当 Channel 关闭后，循环会自动退出。
	for n := range c {
		// 模拟处理数据的时间
		time.Sleep(time.Second)
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

// createWorker 函数用来创建并启动一个 worker。
// 每个 worker 都会启动一个 goroutine 去处理从 channel 中接收到的数据。
// 该函数返回一个 channel 用来接收传入的数据。
func createWorker(id int) chan<- int {
	c := make(chan int) //创建一个无缓冲 Channel。
	go worker(id, c)    //启动 Goroutine，运行 worker 函数，让它监听传入的 Channel。
	return c
}

func main() {
	// 创建两个生产者 generator 函数返回的 channel。
	var c1, c2 = generator(), generator()
	// 创建一个 worker 来消费数据
	var worker = createWorker(0)

	//把生成的数据存起来
	var values []int
	// 初始化计数器 n，存储生成的数据
	n := 0

	// 创建一个定时器，在 10 秒后触发。
	tm := time.After(10 * time.Second)
	// 创建一个定时器，每秒触发一次，用来监控队列的长度。
	tick := time.Tick(time.Second)

	// 无限循环，持续接收生产者的数据并传给消费者，同时处理超时、定时等事件
	for {
		var activeWorker chan<- int
		var activeValue int
		// 如果队列有数据，就选择第一个数据并传给 worker
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1: // 如果从 c1 中接收到数据，将其加入队列
			values = append(values, n)
		case n = <-c2: // 如果从 c2 中接收到数据，将其加入队列
			values = append(values, n)
		case activeWorker <- activeValue: // 将队列中第一个值传给 worker 进行处理
			values = values[1:] // 处理完一个数据后，移除队列中的第一个值

		// 如果超过 500 毫秒没有接收到任何数据，则打印 "time out"
		case <-time.After(500 * time.Millisecond):
			fmt.Println("time out")
		//通过tick,每秒钟检查一次队列的长度，反映当前的系统状态
		case <-tick:
			fmt.Println("queue len =", len(values))

		// 如果 10 秒钟时间到了，就结束程序
		//总的时间去确定程序的运行时长
		case <-tm:
			fmt.Println("bye")
			return

			//default:
			//	fmt.Println("No value received")
		}
	}
}
