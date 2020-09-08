package strings

import (
	"math"
	"strings"
	"sync"
)

var (
	// Base62Codes represents alphabets used to encode baes62
	base62Codes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Base62CodesMap represents alphabets char map with index, used to decode base62 string
	base62CodesMap = make(map[byte]uint64)
	base62Once     = sync.Once{}
)

// EncodeToBase62 will encode number into base62
func EncodeToBase62(n uint64) string {
	if n == 0 {
		return "0"
	}
	result := make([]byte, 0)
	for n > 0 {
		round := n / 62
		remain := n % 62
		result = append(result, base62Codes[remain])
		n = round
	}
	return string(result)
}

// DecodeBase62 will decode base62 string to int
func DecodeBase62(s string) uint64 {
	base62Once.Do(func() {
		if len(base62CodesMap) == 0 {
			for i, v := range base62Codes {
				base62CodesMap[byte(v)] = uint64(i)
			}
		}
	})
	str := strings.TrimSpace(s)
	var result uint64
	for i, ch := range []byte(str) {
		result += base62CodesMap[byte(ch)] * uint64(math.Pow(62, float64(i)))
	}
	return result
}
