package main

import "fmt"

func creatMap() {
	var m map[int]string = make(map[int]string) // 使用make初始化
	var m4 = make(map[int]string, 16)           //make指定长度，避免多次扩容
	fmt.Println(m, len(m))
	fmt.Println(m4, len(m4))
	var m1 map[int]string = map[int]string{ //指定类型
		90: "excellent",
		80: "good",
		60: "pass",
	}
	fmt.Println(m1, len(m1))
	var m2 = map[int]string{ //类型推导
		90: "excellent",
		80: "good",
		60: "pass",
	}
	fmt.Println(m2, len(m2))

}

func operationMap() {
	var fruits = map[string]int{
		"apple":  2,
		"banana": 5,
		"orange": 8,
	}
	// 读取元素
	var score = fruits["apple"]
	fmt.Println(score)
	// 添加或修改
	fruits["pear"] = 3
	fmt.Println(fruits)
	// 删除元素
	delete(fruits, "apple")
	fmt.Println(fruits)
	for k, v := range fruits {
		fmt.Println(k, v)
	}
}

func main() {
	//creatMap()
	operationMap()
}
