package strings

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	// Alphabets represents alphabet list, used for string generation
	Alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GetRandomString will generate a random string with given length
func GetRandomString(n int) string {
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = Alphabets[b%byte(len(Alphabets))]
	}
	return string(bytes)
}

// GetRandomNumber will generate a random string with given length of numbers
func GetRandomNumber(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// SubString will return a sub string with given start position and length
func SubString(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// ContainString will split string and query target string is splited array
func ContainString(source string, sep string, target string) bool {
	for _, s := range strings.Split(source, sep) {
		if s == target {
			return true
		}
	}
	return false
}
