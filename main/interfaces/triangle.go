package interfaces

type Triangle struct {
	Height float64
	Base   float64
}

func (triangle Triangle) getArea() float64 {
	return 0.5 * triangle.Base * triangle.Height
}
