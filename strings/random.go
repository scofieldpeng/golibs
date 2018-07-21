package strings

import (
	"math/rand"
	"time"
)

// Random 生成随机数
func Random(length ...int) string {
	if len(length) == 0 {
		length = make([]int, 1)
		length[0] = 8
	}
	if length[0] < 1 {
		length[0] = 8
	}
	source := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ01234567890"
	rand.Seed(time.Now().UnixNano())
	res := make([]byte, 0, length[0])

	for i := 0; i < length[0]; i++ {
		res = append(res, []byte(source)[rand.Intn(62)])
	}

	return string(res)
}
