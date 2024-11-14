package fib

// 斐波那契数列
// 1, 1, 2, 3, 5, 8, 13, ...
//
//	a, b
//	   a, b
//
// fibonacci 函数返回一个生成斐波那契数列的生成器
// 通过闭包形式，每次调用生成下一个斐波那契数
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
