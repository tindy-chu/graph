package util

func Includes(strSlice []string, str string) bool {
	for _, current := range strSlice {
		if current == str {
			return true
		}
	}
	return false
}
