func repeatedNTimes(A []int) int {
    index := 0
	for ; index < len(A) -2; index ++ {
		if A[index] == A[index + 1] || A[index] == A[index + 2 ]{
			return A[index]
		}
	}
	return A[index + 1]
}