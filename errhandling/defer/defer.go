package main

import (
	"bufio"
	"errors"
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
	//os.O_EXCL 如果文件已存在，会返回一个错误。
	//os.O_CREATE 如果文件不存在，则创建新文件。
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)

	//手动将 err 赋值为一个自定义错误信息 this is a custom error。这只是演示了如何生成一个错误对象。
	err = errors.New("this is a custom error")
	if err != nil {
		//err.(*os.PathError)
		//这里是一个类型断言，检查 err 是否是 *os.PathError 类型。
		//os.PathError 是 Go 中的一个结构体，用于表示文件路径相关的错误。
		//if !ok：
		//如果类型断言失败（即 err 不是 *os.PathError 类型），ok 为 false。
		//这时直接调用 panic(err) 终止程序，并打印错误信息。
		if pathError, ok := err.(*os.PathError); ok {
			// 如果类型断言成功
			// 打印 PathError 的详细信息，包括操作类型（Op）、文件路径（Path）和具体错误（Err）。
			fmt.Printf("%s,%s,%s\n", pathError.Op, pathError.Path, pathError.Err)
		} else {
			// 其他错误直接panic
			panic(err)
		}
		return
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
	writeFile("fib.txt")
}
