package main

import (
	"fmt"
	"math"
)

func main() {
	var num int = 123
	var temp int = 0
	for num != 0 {
		pop := num % 10
		temp = temp*10 + pop
		if temp > math.MaxInt64 || temp < math.MinInt64 {
			return
		}
		num = num / 10
	}
	fmt.Println(temp)
}
