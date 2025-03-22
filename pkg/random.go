package pkg

import (
	"math/rand"
)

const alphanumerics = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a pseudo-random string of length n.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumerics[rand.Int63()%int64(len(alphanumerics))]
	}
	return string(b)
}

// RandomMaxString 大写随机数
func RandomMaxString(n int) string {
	const alphanumerics2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = alphanumerics2[rand.Int63()%int64(len(alphanumerics2))]
	}
	return string(b)
}
