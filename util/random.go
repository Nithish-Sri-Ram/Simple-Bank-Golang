package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// this will generate and return a random integer b/w min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
	// basically the rand.Int63 function returns a random integer between - and n-1
}

// Generate a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// Random owner generates a random owner
func RandomOwner() string {
	return RandomString(6)
}

// To generate random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// This generates a random currency code
func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}

	n := len(currencies)

	return currencies[rand.Intn(n)]
}
