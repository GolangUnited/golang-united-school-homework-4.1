package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	if len(input) == 0 {
		return ""
	}

	myRune := []rune(input)
	smallOutput := ReverseString(string(myRune[1:]))

	output = smallOutput + string(myRune[0])

	return output
}
