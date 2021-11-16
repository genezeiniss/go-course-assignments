package main

import (
	"fmt"
	"strconv"
)

func evenAndOdd(integers []int) {

	for integer := range integers {
		if integer % 2 == 0 {
			fmt.Println(strconv.Itoa(integer) + " is even")
		} else {
			fmt.Println(strconv.Itoa(integer) + " is odd")
		}
	}
}
