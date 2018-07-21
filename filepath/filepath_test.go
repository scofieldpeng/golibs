package filepath

import "testing"

func TestGetRunDir(t *testing.T) {
	path,err := GetRunDir()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log("app run path:",path)
}
