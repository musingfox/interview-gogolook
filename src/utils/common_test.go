package utils

import "testing"

func TestBoolToInt8(t *testing.T) {
	if BoolToInt8(true) == 1 && BoolToInt8(false) == 0 {
		t.Log("success")
	} else {
		t.Log("failed")
	}
}
