//author: richard
package util

import (
	"math/rand"
	"time"
)

//int
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

//string
func RandomString(n int) string {
	var str  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var length = len(str)
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = str[RandomInt(length)]
	}
	return string(buf)
}
