package utils

func Ternary(condition bool, firstVal, secondVal int) int {
	if condition {
		return firstVal
	} else {
		return secondVal
	}
}
