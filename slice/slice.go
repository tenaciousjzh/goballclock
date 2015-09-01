package slice

func Same(s1 []int, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func Shift(s []int) (int, []int) {
	val, s := s[0], s[1:]
	return val, s
}

func Push(s []int, val int) []int {
	return append(s, val)
}

func Pop(s []int) (int, []int) {
	val, s := s[len(s)-1], s[:len(s)-1]
	return val, s
}
