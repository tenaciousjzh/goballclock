package slice

import (
	"testing"
)

func TestSame(t *testing.T) {
	s1 := []int{0, 1, 2}
	s2 := []int{0, 1, 2}

	same := Same(s1, s2)
	if same == false {
		t.Errorf("s1 and s2 should be the same")
	}
}

func TestNotSame(t *testing.T) {
	s1 := []int{0, 1, 2}
	s2 := []int{0, 4, 3}

	same := Same(s1, s2)
	if same == true {
		t.Errorf("s1 and s2 shoudl not be the same")
	}
}

func TestShift(t *testing.T) {
	s1 := []int{7, 8, 9}
	expected := []int{8, 9}
	head, tail := Shift(s1)

	if head != 7 {
		t.Errorf("Shift should return the front of the slice: 7")
	}

	same := Same(expected, tail)

	if same == false {
		t.Errorf("shift should have returned the tail of the slice")
	}
}

func TestPush(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := Push(s1, 4)

	if len(s2) != 4 {
		t.Errorf("Push should have added an element to the slice")
	}

	expected := []int{1, 2, 3, 4}
	same := Same(expected, s2)

	if same == false {
		t.Errorf("Push should have added 4 to the end of the slice")
	}
}

func TestPop(t *testing.T) {
	s1 := []int{1, 2, 3}
	last, front := Pop(s1)

	if last != 3 {
		t.Errorf("Pop should have returned 3 from the slice s1")
	}

	expected := []int{1, 2}
	same := Same(expected, front)
	if same == false {
		t.Errorf("Pop should have removed the last element from the slice ")
	}
}
