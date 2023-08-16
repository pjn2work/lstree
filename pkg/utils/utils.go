package utils

func FillWith(char string, n int) string {
	res := ""
	for ; n > 0; n-- {
		res += char
	}
	return res
}

func GetStringLen(str string) int {
	r := []rune(str)
	return len(r)
}
