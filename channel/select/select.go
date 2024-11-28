package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

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

func createWorker(id int) chan<- int {
	c := make(chan int) //创建一个无缓冲 Channel。
	go worker(id, c)    //启动 Goroutine，运行 worker 函数，让它监听传入的 Channel。
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
			//default:
			//	fmt.Println("No value received")
		}
	}
}
