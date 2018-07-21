package number

import "testing"

func TestRandom(t *testing.T) {
	str1 := Random()
	if len(str1) != 4 {
		t.Error("string need to length 4,get ", len(str1))
	}

	str2 := Random(5)
	if len(str2) != 5 {
		t.Errorf("string need to length 5,get ", len(str2))
	}

	str3 := Random(33)
	if len(str3) != 4 {
		t.Errorf("string need to length 4,get: ", len(str3))
	}

	str4 := Random(-1)
	if len(str4) != 4 {
		t.Errorf("string need to length 4,get: ", len(str4))
	}
}
