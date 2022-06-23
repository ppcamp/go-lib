package random

import (
	"math/rand"
	"strings"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321-!_#@*"
const defaultLen = 30
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// String generate a rand string with `n` size, basing on default chars
func String(n int) string {
	return randString([]byte(letterBytes), n)
}

// String generate a rand string with `n` size, basing on str param
func StringFrom(str []byte, n int) string {
	if n <= 0 {
		n = defaultLen
	}

	if len(str) == 0 {
		str = []byte(letterBytes)
	}

	return randString(str, n)
}

// randString was built upon this https://stackoverflow.com/a/31832326/10013122
func randString(letters []byte, n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			sb.WriteByte(letters[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
