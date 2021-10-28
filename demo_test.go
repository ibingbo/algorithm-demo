/*
 * Copyright(C) 2021 Baidu Inc. All Rights Reserved.
 * Author: zhangbingbing
 * Date: 2021/10/27
 * File: demo_test.go
 */

/*
 * DESCRIPTION
 *   TODO file function desc
 */

// Package src TODO package function desc
package src

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestReverseNumber(t *testing.T) {
	pp := []struct {
		input  int
		expect int
	}{
		{123, 321},
		{1234, 4321},
	}

	t.Run("reverse number", func(t *testing.T) {
		for _, p := range pp {
			if p.expect != ReverseNumber(p.input) {
				t.Errorf("error")
			}
		}
	})
}

func TestReverseNumber2(t *testing.T) {
	assert := assert.New(t)
	pp := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{1234, 4321},
	}
	for _, p := range pp {
		assert.Equal(p.expected, ReverseNumber(p.input))
	}

}

func TestReverseNumber3(t *testing.T) {
	convey.Convey("reverse number", t, func() {
		pp := []struct {
			input    int
			expected int
		}{
			{123, 321},
			{1234, 4321},
		}
		for _, p := range pp {
			convey.So(p.expected, convey.ShouldEqual, ReverseNumber(p.input))
		}
	})

	convey.Convey("reverse number 1", t, func() {
		applay := gomonkey.ApplyFunc(ReverseNumber, func(a int) int {
			return a
		})
		defer applay.Reset()
		pp := []struct {
			input    int
			expected int
		}{
			{123, 321},
			{1234, 4321},
		}
		for _, p := range pp {
			convey.So(p.expected, convey.ShouldEqual, ReverseNumber(p.input))
		}
	})
}

func TestReverseNumber4(t *testing.T) {
	monkey.Patch(ReverseNumber, func(a int) int {
		return a
	})
	assert := assert.New(t)
	pp := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{1234, 4321},
	}
	for _, p := range pp {
		assert.Equal(p.expected, ReverseNumber(p.input))
	}
	monkey.UnpatchAll()

}

func TestReverseNumber5(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	})
	defer monkey.UnpatchAll()
	t.Log(time.Now())
}
