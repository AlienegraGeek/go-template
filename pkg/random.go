package pkg

import (
	"math/rand"
	"strconv"
	"time"
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

const as = "136145481759273241471451054914857641823742613747"

func RandomFloat64(n int) float64 {
	b := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		if i == 2 {
			b[i] = '0'
		} else if i == 1 {
			b[i] = '.'
		} else {
			b[i] = as[(rand.Intn(len(as)))]
		}

	}
	float, _ := strconv.ParseFloat(string(b), 4)
	return float
}
