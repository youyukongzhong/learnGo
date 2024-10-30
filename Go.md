# Go开发工程师
## 基础知识
***
### 内建容器
#### > 3-1 数组(array)
```
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
```

> * 数组是值类型  
>   * [10]int 和 [20]int是不同类型
>   * 调用func f(arr [10]int)会拷贝数组
>   * 在go语言中一般不直接使用数组,而是用<u>**切片**</u>
> * 参数传递是<u>**值传递**</u>

### Slice(切片)
* 切片是数组的视图  
```
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
}
```

