package loop

import (
	"fmt"
	"strconv"
)

func IsEven(number int) bool {
	if number%2 == 0 {
		fmt.Println(strconv.Itoa(number) + " is even")
		return true
	} else {
		fmt.Println(strconv.Itoa(number) + " is odd")
		return false
	}
}
