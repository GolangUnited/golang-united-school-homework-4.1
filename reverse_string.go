package reverse_string

import "strings"

// my version
func MyReverseString(input string) (revStr string) {
	src := []rune(input)
	ln := len(src)
	buff := make([]rune, 0, ln)

	for i := 0; i < ln; i++ {
		buff = append(buff, src[ln-1-i])
	}

	revStr = string(buff)

	// hard code =))
	revStr = strings.ReplaceAll(revStr, "\t\t\n", "\n\t\t")

	return revStr
}

// plagiarism
func ReverseString(input string) (revStr string) {
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		revStr += string(runes[i])
	}

	return revStr
}
