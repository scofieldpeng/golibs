package strings

import "testing"

func TestRandom(t *testing.T) {
	str1 := Random()
	if len(str1) != 8 {
		t.Error("want get length 8,get: ", len(str1))
	}
	str2 := Random(16)
	if len(str2) != 16 {
		t.Error("want get length 16,get: ", len(str2))
	}
	str3 := Random(-1)
	if len(str3) != 8 {
		t.Error("want get length 8,get:", len(str3))
	}
}
