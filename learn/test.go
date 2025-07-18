package main

import "fmt"

type Geometry interface {
	Perimeter() int
}
type Square struct {
	a int
	b int
}

func (s Square) Perimeter() int {
	return s.a + s.b
}
func main() {
	fmt.Println("hello, world")
	var s Geometry = Square{5, 10}
	v, ok := s.(Square)
	fmt.Println(v, ok)
}
