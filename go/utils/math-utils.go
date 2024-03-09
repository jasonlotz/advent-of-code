package utils

func Abs(int int) int {
	if int < 0 {
		return -int
	}
	return int
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
