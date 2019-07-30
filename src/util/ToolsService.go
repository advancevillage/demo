//author: richard
package util

import (
	"math/rand"
	"time"
)

func RandomInt(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func MinInt(x, y int) int {
	if x >= y {
		return y
	} else {
		return x
	}
}

func MaxInt(x, y int) int {
	if x >= y {
		return  x
	} else {
		return y
	}
}
