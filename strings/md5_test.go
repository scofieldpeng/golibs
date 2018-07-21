package strings

import "testing"

func TestMD5(t *testing.T) {
	str1 := ""
	if MD5(str1) != "" {
		t.Error("emtpy md5 should empty")
	}
	str2 := "123456"
	if MD5(str2) != "e10adc3949ba59abbe56e057f20f883e" {
		t.Error("md5('123456') get wrong hash value")
	}
}
