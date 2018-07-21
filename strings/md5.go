package strings

import (
	"crypto/md5"
	"fmt"
)

// MD5 生成MD5
func MD5(data string) string {
	if data == "" {
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}
