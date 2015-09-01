package slice

//Same provides a comparison capability to determine if
//two slices have the same elements in the same order.
//This provides the condition for stopping the evaluation
//in the ballClock object when the queue of balls
//returns to the original order.
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

//Shift returns the first element from s and the tail of s
func Shift(s []int) (int, []int) {
	val, s := s[0], s[1:]
	return val, s
}

//Push appends a value to the end of s
func Push(s []int, val int) []int {
	return append(s, val)
}

//Pop removes the last element from s and returns it along with the rest of s
func Pop(s []int) (int, []int) {
	val, s := s[len(s)-1], s[:len(s)-1]
	return val, s
}
