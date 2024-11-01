package main

import "fmt"

func main() {
	//创建map
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

	//遍历map
	//map中的k,v对是无序的,需要手动使用slice对map的key进行排序
	fmt.Println("--------Traversing map--------")
	for k, v := range m {
		fmt.Println(k, v)
	}

	//获取map中的value
	fmt.Println("--------Getting values--------")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	//对于 nil 的 map，读取一个不存在的键不会产生错误，只会返回该值类型的零值
	//判断map的值存不存在
	//ok 是一个布尔值，用于表示在 map 中是否存在指定的键
	if courName, ok := m["crse"]; ok {
		fmt.Println(courName)
	} else {
		fmt.Println("key does not exist")
	}

	//删除map中的值
	fmt.Println("--------Deleting map--------")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
