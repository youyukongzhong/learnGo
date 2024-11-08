package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes我爱GO语言" // UTF-8编码
	//fmt.Printf("%s\n", []byte(s))
	//fmt.Println(s)
	// 1. 将字符串 s 转换为字节切片,并输出每个字节的十六进制值。
	// []byte(s) 将字符串 s 转换成字节数组,以便可以按原始字节数据读取。
	//输出结果:596573e68891e788b1474fe8afade8a880
	//英文字符 “Yes” 用 ASCII 编码，占一个字节。
	//中文字符 “我爱GO语言” 用 UTF-8 编码，每个汉字占 3 字节。
	fmt.Printf("%x\n", []byte(s))

	// 2.逐个打印字符串的每个字节的十六进制值,逐个输出 s 转换成字节切片后每个字节的编码。
	// 输出结果: 59 65 73 e6 88 91 e7 88 b1 47 4f e8 af ad e8 a8 80
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()

	// 3.使用 for range 遍历字符串 s 中的每个字符
	//输出中每个汉字字符用一个 rune 表示，英文字符依旧用一个字节表示。
	// 输出结果: (0 59) (1 65) (2 73) (3 6211) (6 7231) (9 47) (10 4f) (11 8bed) (14 8a00)
	// 进行了UTF-8的解码
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %x) ", i, ch)
	}
	fmt.Println()

	// 4.输出字符串 s 的 符文（rune）数量。中文字符被视作单个字符，即使其 UTF-8 编码占用多个字节。
	//Rune count: 符文计数
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// 5.通过 utf8.DecodeRune 解码字符串的字节数组（[]byte），这段代码会逐个读取字节，并将其转换成一个 rune(字符)。
	// (utf8.DecodeRune：将"字节"解码为"字符")
	// 每次解码得到一个字符 ch 和其占用的字节数 size，然后从 bytes 切片中跳过这些字节。最终输出每个字符 ch，与原始字符串内容一致。
	bytes := []byte(s) //先将s转为字节,因为DecodeRune方法接受的是byte类型
	for len(bytes) > 0 {
		//func DecodeRune(p []byte) (r rune, size int)
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 6.这里将 s 转换成 rune 切片，然后遍历每个字符（以 rune 表示）。
	// []rune(s)：将字符串转换为 rune 数组，遍历时每个字符都会按 Unicode 编码处理。
	// 这样可以确保遍历时每个字符都被正确处理，即中文字符也只会被视为一个字符。
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	//strings包里有其它关于string的操作
}
