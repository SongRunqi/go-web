package main

import "fmt"

// Geometry 是一个接口，声明了一个方法Perimeter
type Geometry interface {
	Perimeter() int
}

// Square 是一个类，实现了Geometry这个接口，因为他实现了Geometry的所有方法
type Square struct {
	a int
	b int
}

// Perimeter 注意关键字func和方法名Perimeter之间的(s Square)，表明这个方法是Square类的方法
func (s Square) Perimeter() int {
	return s.a + s.b
}

// end Square

func main() {
	fmt.Println("hello, go,wnd！")
	var s Geometry = Square{5, 10}
	v, ok := s.(Square)
	fmt.Println(v, ok)
}
