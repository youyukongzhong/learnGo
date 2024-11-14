package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertToBin(n int) string {
	result := ""
	//将整数转成字符串(10  -> "10")必须使用strconv.Itoa(num)来转换
	//使用string(num)会将转换成num在uft-8对应的字符, 会出现乱码现象.
	//lsb:最低有效位
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		//panic 报错
		panic(err)
	}

	printFileContents(file)

}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	for {
		fmt.Println("abc")
	}
}

func main() {
	//fmt.Fscanf()
	//fmt.Fprintf()
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("basic/branch/abc.txt")

	//像文件一样打印字符串
	s := `abc"d"
	kkkk
	a123
	
	p`
	printFileContents(strings.NewReader(s))
	//forever()
}
