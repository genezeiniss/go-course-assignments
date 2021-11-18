package interfaces

import "fmt"

type shape interface {
	getArea() float64
}

func PrintArea(s shape) {
	fmt.Println(s.getArea())
}
