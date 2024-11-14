package main

import (
	"bufio"
	"fmt"
	"learngo_muke/functional/fib"
	"os"
)

// defer 运行逻辑为先进后出
func tryDefer() {
	// 体现了 defer 在语句时计算
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file) // 写文件
	defer writer.Flush()            // 程序结束后 导入文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
