package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	if len(input) == 0 {
		return ""
	}

	smallOutput := ReverseString(input[1:])
	myRune := []rune(input)
	output = smallOutput + string(myRune[0])

	return output
}
