package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}

}

func div(a, b int) (q, r int) {
	return a / b, r * b
}

func apply(op func(int, int) int, a, b int) int {
	//使用 reflect包中的 ValueOf() 方法获取 op 函数的值,并通过 Pointer() 方法获取该函数的指针。这是为了能够找到函数在内存中的地址。
	p := reflect.ValueOf(op).Pointer()
	//使用 runtime 包中的 FuncForPC() 函数，通过指针 p 获取对应的函数对象,并调用 Name() 方法获取该函数的名称（字符串形式）。
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//func pow(a, b int) int {
//	return int(math.Pow(float64(a), float64(b)))
//}

// 可变参数列表, ...int表示可传入多个int
// 求和
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(3, 3, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 4, 54))

	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)

}
