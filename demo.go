/*
 * Copyright(C) 2021 Baidu Inc. All Rights Reserved.
 * Author: zhangbingbing
 * Date: 2021/10/27
 * File: demo.go
 */

/*
 * DESCRIPTION
 *   TODO file function desc
 */

// Package src TODO package function desc
package main

import (
	"math"
)

func ReverseNumber(num int) int {
	temp := 0
	for num != 0 {
		pop := num % 10
		temp = temp*10 + pop
		if temp > math.MaxInt64 || temp < math.MinInt64 {
			return 0
		}
		num = num / 10
	}
	return temp
}
