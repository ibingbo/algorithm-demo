/*
 * Author: zhangbingbing
 * Date: 2021/11/03
 * File: bitmap.go
 */

/*
 * DESCRIPTION
 *   TODO file function desc
 */

// Package src TODO package function desc
package src

import (
	"testing"
)

func TestBitMap_Add(t *testing.T) {
	bitMap := NewBitMap(10)
	bitMap.Add(7)
	t.Log(bitMap.bits)
}

func TestBitMap_Set(t *testing.T) {
	bitMap := NewBitMap(10)
	bitMap.Set(7)
	t.Log(bitMap.bits)
}

func TestBitMap_Get(t *testing.T) {
	bitMap := NewBitMap(10)
	bitMap.Set(8)
	t.Log(bitMap.Exist(8))
	t.Log(bitMap.Get(8))
	bitMap.Unset(8)
	t.Log(bitMap.Exist(8))
	t.Log(bitMap.Get(8))
}
