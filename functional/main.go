package main

import (
	"bufio"
	"fmt"
	"io"
	"learngo_muke/functional/fib"
	"strings"
)

// intGen 是一个函数类型，返回一个整数
type intGen func() int

// Read 实现了 io.Reader 接口
// 它将斐波那契数转换为字符串形式，并写入到提供的字节切片 p 中
func (g intGen) Read(p []byte) (n int, err error) {
	// 获取下一个斐波那契数
	next := g()
	//if next > 10000 {
	//	// 如果超过 10000，则返回 EOF 表示数据结束
	//	return 0, io.EOF
	//}
	// 将斐波那契数转换为字符串形式，并附加换行符
	s := fmt.Sprintf("%d\n", next)

	// 将字符串写入到字节切片 p 中
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

// printFileContents 可以逐行打印文件内容的函数
// 处理任何支持“逐行读取”的数据源，比如文件、字符串、甚至网络数据。
func printFileContents(reader io.Reader) {
	// NewScanner  扫描器（Scanner），可以逐行读取数据
	scanner := bufio.NewScanner(reader)

	//一个循环，逐行读取 reader 中的内容，并打印出来。
	for i := 0; i < 15 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
}

func main() {
	// 创建一个斐波那契数列生成器
	var f intGen = fib.Fibonacci()

	// 打印斐波那契数列内容
	printFileContents(f)
}
