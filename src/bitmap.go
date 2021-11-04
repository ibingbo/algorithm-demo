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
	"bytes"
	"encoding/binary"
)

type BitMap struct {
	bits []byte
	max  int
}

func NewBitMap(max int) *BitMap {
	bits := make([]byte, max/8+1)
	return &BitMap{
		bits: bits,
		max:  max,
	}
}
func (bitMap *BitMap) Add(num int) {
	idx := num / 8
	pos := num % 8
	bitMap.bits[idx] |= 1 << pos
}

func (bitMap *BitMap) Remove(num int) {
	idx := num / 8
	pos := num % 8
	bitMap.bits[idx] &= ^(1 << pos)
}

func (bitMap *BitMap) Set(num int) {
	idx := num >> 3
	pos := num & 0x07
	bitMap.bits[idx] |= 1 << pos
}

func (bitMap *BitMap) Unset(num int) {
	idx := num >> 3
	pos := num & 0x07
	bitMap.bits[idx] &= ^(1 << pos)
}

func (bitMap *BitMap) Get(num int) int {
	idx := num / 8
	pos := num % 8
	res, _ := binary.ReadVarint(bytes.NewBuffer(bitMap.bits[idx : idx+1]))
	return int(res) & (1 << pos)
}

func (bitMap *BitMap) Exist(num int) bool {
	idx := num / 8
	pos := num % 8
	return bitMap.bits[idx]&(1<<pos) != 0
}
