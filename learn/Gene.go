package main

import "fmt"

// 定义自定义接口
type MyInterface interface {
	String() string
	Compare(other MyInterface) bool
}

type MyA struct {
	a int
}

func (m MyA) String() string {
	return "MyA"

}

func (m MyA) Compare(other MyInterface) bool {
	return m.a == other.(MyA).a
}

// 使用自定义接口作为泛型约束
func Contains[T MyInterface](arr []T, x T) bool {
	for _, v := range arr {
		if v.Compare(x) { // 使用接口方法进行比较
			return true
		}
	}
	return false
}

//	func Contains[T comparable](arr []T, x T) bool {
//		for _, v := range arr {
//			if v == x {
//				return true
//			}
//		}
//		return false
//	}
func main() {
	//intArr := []int{1, 2, 3, 4, 5}
	//threeIsExits := Contains(intArr, 3)
	//fmt.Println(threeIsExits)
	//
	//stringArr := []string{"a", "b", "c"}
	//threeIsExits = Contains(stringArr, "c")
	//fmt.Println(threeIsExits)
	var t1 MyInterface = MyA{1}
	list := []MyInterface{t1, MyA{2}}
	fmt.Println(Contains(list, t1))
}
