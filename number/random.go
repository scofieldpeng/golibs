package number

import (
	"math/rand"
	"time"
)

// Random 随机数,默认返回4位随机数,最大
func Random(length ...int) string {
	if len(length) == 0 {
		length = make([]int, 1)
		length[0] = 4
	}
	if length[0] < 1 || length[0] > 32 {
		length[0] = 4
	}
	source := "0123456789"
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, 0, length[0])
	for i := 0; i < length[0]; i++ {
		res = append(res, []byte(source)[rand.Intn(10)])
	}

	return string(res)
}
