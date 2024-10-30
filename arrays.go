package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	//数组的多种定义方法
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}

	//二维数组
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	//遍历数组
	/*
		for i := 0; i < len(arr3); i++ {
			fmt.Println(arr3[i])
		}
	*/
	//遍历数组 range 关键字可以获取数组的下标
	for i := range arr3 {
		fmt.Println(arr3[i])
	}

	//获取下标和值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 只要值, 不需要下标
	for _, v := range arr3 {
		fmt.Println(v)
	}

	fmt.Println("printArray(arr1")
	printArray(arr1)

	fmt.Println("printArray(arr3")
	printArray(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
