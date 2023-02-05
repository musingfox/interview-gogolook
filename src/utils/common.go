package utils

func BoolToInt8(b bool) int8 {
	var bitSetVar int8
	if b {
		bitSetVar = 1
	}

	return bitSetVar
}
