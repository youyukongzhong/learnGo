package main

import (
	"fmt"
	"time"
)

// worker :用于从 Channel 中接收数据并处理
func worker(id int, c chan int) {
	//遍历 Channel，直到 Channel 被关闭。
	//每次接收一个数据 n，输出到控制台。
	//range 的作用：优雅地处理 Channel 的关闭，当 Channel 关闭后，循环会自动退出。
	for n := range c {
		//n, ok := <-c
		//if !ok {
		//	break
		//}
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

// createWorker :这是一个工厂函数，用于创建每个 worker，并返回用于发送数据的 Channel。
// chan<- int ：表示返回的 Channel 只能用于发送数据，限制了 Channel 的使用方向，增加了代码的安全性。
func createWorker(id int) chan<- int {
	c := make(chan int) //创建一个无缓冲 Channel。
	go worker(id, c)    //启动 Goroutine，运行 worker 函数，让它监听传入的 Channel。
	return c
}

// chanDemo :演示了如何使用多个 Channel 与多个 Goroutine 通信。
func chanDemo() {
	// 创建了一个 channels 数组，用于存储 10 个发送 Channel。
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		// 每个 Channel 对应一个 Worker Goroutine。
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

// bufferedChannel 演示了带缓冲的 Channel。
func bufferedChannel() {
	c := make(chan int, 3) //创建一个缓冲区大小为 3 的 Channel。
	go worker(0, c)        // 启动一个 worker goroutine 来接收数据并处理

	// 依次发送 4 个数据
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' // 这一行会阻塞，直到有接收者消费数据
	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	// 关闭 channel
	close(c)

	// 向已关闭的 channel 发送数据
	// 使用 select 来避免 panic
	//go func() {
	//	select {
	//	case c <- 'e':
	//	default:
	//		fmt.Println("channel is close, cannot send data.")
	//	}
	//}()

	// 使用 defer(推迟) 和 recover 来捕获 panic
	// recover 必须在 defer 语句中使用，且只能在发生 panic 后的 defer 中被调用。
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	c <- 'e'

	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("channel as first-class citizen")
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
}
