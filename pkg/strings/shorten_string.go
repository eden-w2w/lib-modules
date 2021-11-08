package strings

func ShortenString(str string, length int, suffix string) string {
	runes := []rune(str)
	if len(runes) <= length {
		return str
	}
	str = string(runes[:length])
	return str + suffix
}
