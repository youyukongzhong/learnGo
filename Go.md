# Go开发工程师
## 基础知识
***
### 内建容器
#### array (数组)


知识点:
> * 数组是值类型  
> * [10]int 和 [20]int是不同类型
> * 调用func f(arr [10]int)会拷贝数组
> * 在go语言中一般不直接使用数组,而是用<u>**切片**</u>
> * 参数传递是<u>**值传递**</u>

案例:
``` go
// 数组的定义 
// 遍历数组
// 获取下标和值
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

#### Slice(切片)
* Slice(切片)是数组的视图
##### Slice的扩展  
<img src="image-1.png" alt="Slice扩展" style="zoom:31%;" />
<img src="image-2.png" alt="Slice扩展" style="zoom:39%;" />
<img src="image-3.png" alt="Slice的实现" style="zoom:36%;" />

##### Slice的 <u>添加元素</u> 操作

> * 向Slice添加元素,添加元素时如果超越cap，系统会重新分配更大的底层数组
> * 由于值传递的关系，必须接收append的返回值

案例:

```go
package main
import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	s1 := arr[2:]
	fmt.Println("s1 = ", s1)
	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice")
	fmt.Println(s2)
	s2 = s2[:5]
	fmt.Println(s2)
	s2 = s2[2:]
	fmt.Println(s2)

	// slice的扩展
	fmt.Println("now arr: ", arr)
	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	fmt.Println("now arr: ", arr)
	s1 = arr[2:6]
	s2 = s1[3:5] // [s1[3], s1[4]]
	//fmt.Println("s1: ", s1) //s1:  [2 3 4 5]
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n",
		s2, len(s2), cap(s2))
	//fmt.Println("s2: ", s2) //s2:  [5 6]

	//slice的添加(append)
	//s = append(s, val)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	// s4 and s5 no longer view arr
	fmt.Println("arr = ", arr)

}
```
##### Slice<u>**其它操作**</u>
> * 创建slice
> * 拷贝
> * 删除
> * 去掉开头值 和 结束值

案例:
```go
package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
	//创建slice
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)
    
	//拷贝操作
	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)

	//删除操作(用append函数,"..."三个点表示后面所有的值)
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	//去掉开头值
	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	//去掉结束值
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]

	fmt.Println(tail)
	printSlice(s2)
}
```
#### Map
##### Map结构定义
> * map[k]v
> * map[k1]map[k2]v

##### Map的操作  
> * 创建:make(map[string]int)  
> * 获取元素: m[key]  
> * key不存在时,获得Value类型的初始值  
> * 用value, ok := m[key] 来判断是否存在key
> * 用delete函数删除一个key

##### Map的遍历
> * 使用range遍历key, 或者遍历key,value对
> * 不保证遍历顺序,如需顺序,需手动对key排序
> * 使用len获得元素的个数

##### Map的key
> * map使用哈希表，必须可以比较相等
> * 除了slice, map,function的内建类型都可以作为key
> * Struct类型不包含上述字段，也可作为key

案例:
```go
func main() {
	//构造
	fmt.Println("--------Creating map--------")
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golong",
		"site":    "imooc",
		"quality": "notbad",
	}

	//make定义法:使用 make 创建的 map 是已初始化的，可以直接使用。
	m2 := make(map[string]int) //m2 == empty map

	//var定义法:使用 var 声明的 map 是 nil，未初始化，可以读取 nil map 中的值，返回该类型的零值,但不能直接用于存储键值对，必须用 make 初始化后才能使用。
	var m3 map[string]int // m3 == nil

	fmt.Printf("m1 = %v\nm2 = %v\nm3 = %v\n", m, m2, m3)

	//遍历
	//map中的k,v对是无序的,需要手动使用slice对map的key进行排序
	fmt.Println("--------Traversing map--------")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//取值
	fmt.Println("--------Getting values--------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//判断map的值存不存在
	//对于 nil 的 map，读取一个不存在的键不会产生错误，只会返回该值类型的零值
	//ok 是一个布尔值，用于表示在 map 中是否存在指定的键
	if courName, ok := m["crse"]; ok {
		fmt.Println(courName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除
	fmt.Println("--------Deleting map--------")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
```



