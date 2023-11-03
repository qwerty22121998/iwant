package iwant

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	LowerAlpha = "abcdefghijklmnopqrstuvwxyz"
	UpperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Number     = "0123456789"
	Symbol     = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

// RandString return random string in pool with length n (https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go)
func RandString(n int, pool string) string {
	idxBits := int64(math.Ceil(math.Log2(float64(len(pool)))))
	idxMask := int64(1<<idxBits - 1)
	idxMax := 63 / idxBits
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), idxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), idxMax
		}
		if idx := int(cache & idxMask); idx < len(pool) {
			b[i] = pool[idx]
			i--
		}
		cache >>= idxBits
		remain--
	}
	return string(b)
}
