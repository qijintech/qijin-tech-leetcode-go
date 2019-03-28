func numJewelsInStones(J string, S string) int {
	res := 0
	jewels := make(map[rune] bool,len(J))
	for _, value := range J {
		jewels[value] = true
	}
	for _, value := range S {
		if jewels[value] {
			res ++
		}
	}
	return res

}