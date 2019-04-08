func toLowerCase(str string) string {
	b := make([]rune,len(str))
	offset := 'a' - 'A'
	for index, value := range str {
		if value >= 'A' && str[index] <='Z'{
			b[index] = value + offset
		}else {
			b[index] = value
		}
	}
	return string(b)
}