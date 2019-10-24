package strings

import "crypto/rand"

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