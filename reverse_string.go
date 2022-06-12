package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	rune := []rune(input)
	for i, j := 0, len(rune)-1; i < j; i, j = i+1, j-1 {
		rune[i], rune[j] = rune[j], rune[i]
	}
	return string(rune)
}
