package main

import "testing"

func TestEvenOrOdd(t *testing.T)  {

	var integers []int

	for i := 0; i <= 10; i++ {
		integers = append(integers, i)
	}

	evenAndOdd(integers)
}
