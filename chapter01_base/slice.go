package main

import "fmt"

func creatSlice() {
	// 使用make创建
	var s1 []int = make([]int, 5, 8)
	s2 := make([]int, 8) // 满容切片
	var s3 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3, len(s3), cap(s3))
	// 赋值
	s3[0] = 0
	fmt.Println(s3)

}

func appendSlice() {
	/*
			每次append都会形成新的切片变量，如果底层数组没有扩容，那么就共享数据，
		如果扩容，则数据不共享
	*/
	var s1 []int = []int{1, 2, 3, 4, 5}
	var s2 = append(s1, 6)
	s1[0] = 10
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	var s3 []int = make([]int, 5, 8)
	s4 := append(s3, 6)
	s3[0] = 10
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))

}

func copySlice() {
	/*
		copy 返回copy的个数
	*/
	var s = make([]int, 5, 8)
	for i := 0; i < len(s); i++ {
		s[i] = i + 1
	}
	var s1 = make([]int, 2, 6)
	fmt.Println(s, s1)
	var s2 = copy(s1, s)
	fmt.Println(s2)
	fmt.Println(s, s1)
}
func main() {
	//creatSlice()
	//appendSlice()
	copySlice()
}
