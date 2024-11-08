package main

import "fmt"

// 1. 函数定义
func lengthOfNonRepeatingSubStr(s string) int {
	// 2. 变量初始化
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	// 3. 遍历字符串
	//将字符串 s 转换为 byte 切片后遍历，每个字符以 ch 表示，其索引为 i。
	//之所以将 s 转换为 byte 切片，是因为代码仅处理 ASCII 字符。
	//如果想支持 Unicode 字符，可以将 byte 换成 rune
	for i, ch := range []rune(s) {
		//4. 判断是否出现重复字符
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		// 5. 更新最长子串长度
		//i - start + 1 表示 start 到 i 之间的字符数
		//这就是从 start 到 i 之间的字符总数。
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		// 6. 更新字符最近出现的位置
		lastOccurred[ch] = i
	}
	// 7. 返回结果
	return maxLength
}
func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepeatingSubStr(""))
	fmt.Println(lengthOfNonRepeatingSubStr("b"))
	fmt.Println(lengthOfNonRepeatingSubStr("abcdef"))
	fmt.Println(lengthOfNonRepeatingSubStr("让我们说中文"))
	fmt.Println(lengthOfNonRepeatingSubStr("一二三三二一"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
