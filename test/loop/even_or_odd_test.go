package loop

import (
	"go-course-assignment/main/loop"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {

	var numbers []int

	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for number := range numbers {
		loop.IsEven(number)
	}
}
