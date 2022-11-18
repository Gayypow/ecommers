package converter

import (
	"strconv"
)

// Check if string in slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Converting uint to string
func UintToStr(num uint) string {
	return strconv.FormatUint(uint64(num), 10)
}

func StrToInt(num string) (int, error) {
	newInt, err := strconv.Atoi(num)
	return newInt, err
}

func IntToStr(num int) string {
	return strconv.Itoa(num)
}

func StrToUint(num string) uint {
	newInt, _ := strconv.ParseUint(num, 10, 32)
	return uint(newInt)
}
