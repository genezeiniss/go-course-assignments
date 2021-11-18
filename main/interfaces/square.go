package interfaces

type Square struct {
	SideLength float64
}

func (square Square) getArea() float64 {
	return square.SideLength * square.SideLength
}
