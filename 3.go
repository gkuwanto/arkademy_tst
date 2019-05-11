package main

import (
	"fmt"
	"math/rand"
)

func randString() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func notIn(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}

func randXS(n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		r := randString()
		if notIn(r, result) {
			result = append(result, r)
		} else {
			i--
		}
	}
	return result
}
func cetak(n int) {
	res := randXS(n)
	for i := range res {
		fmt.Println(res[i])
	}
}

func main() {
	cetak(2)
}
