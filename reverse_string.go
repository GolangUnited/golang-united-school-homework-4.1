package reverse_string

func ReverseString(input string) string {
	ln := len(input)
	buff := make([]rune, ln)

	for i, c := range input {
		buff[ln-1-i] = c
	}

	return string(buff)
}
