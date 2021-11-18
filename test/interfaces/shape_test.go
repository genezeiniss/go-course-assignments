package interfaces

import (
	"go-course-assignment/main/interfaces"
	_ "go-course-assignment/main/interfaces"
	"testing"
)

func TestShape(t *testing.T) {

	triangle := interfaces.Triangle{Height: 15, Base: 10}
	interfaces.PrintArea(triangle)

	square := interfaces.Square{SideLength: 17}
	interfaces.PrintArea(square)
}
