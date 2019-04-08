func findTheDifference(s string, t string) byte {
	var res byte
	for _, value := range s {
		res ^= byte(value)
	}
	for _, value := range t {
		res ^= byte(value)
	}
	return res
}