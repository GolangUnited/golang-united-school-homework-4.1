package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	sb := []rune(input)
	for i := len(sb) - 1; i >= len(sb)/2; i-- {
		b := len(sb) - 1 - i
		sb[i], sb[b] = sb[b], sb[i]
	}
	output = string(sb)
	return output
}
