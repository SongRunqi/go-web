package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到panic:", r)
		}
	}()

	arr := []int{1, 2, 3}
	fmt.Println("访问正常索引:", arr[1])
	go fmt.Println("this is a go routin")
	// 这里会触发panic
	fmt.Println("访问越界索引:", arr[5])

	// 这行代码不会执行
	fmt.Println("这行不会被打印")
}
