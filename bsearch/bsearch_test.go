package bsearch

import (
	"testing"
)

func TestMultipeValues(t *testing.T) {
	t.Log("Testing miltiple values")
	array := []int{0, 0, 0, 0, 0, 0, 0}
	result := BSearch(array, 0)
	if result != 3 {
		t.Error("Incorrect result")
	}
}

func TestSingleValue(t *testing.T) {
	t.Log("Testing single value")
	array := []int{1, 2, 3, 4, 5, 6, 7}
	result := BSearch(array, 2)
	if result != 1 {
		t.Error("Incorrect result")
	}
}

func TestEmptyArray(t *testing.T) {
	t.Log("Testing empty array")
	array := []int{}
	result := BSearch(array, 1)
	if result != -1 {
		t.Error("Incorrect result")
	}
}

func TestFirstInArray(t *testing.T) {
	t.Log("Testing first element in array")
	array := []int{1, 2, 3, 4, 5, 6, 7}
	result := BSearch(array, 1)
	if result != 0 {
		t.Error("Incorrect result")
	}
}

func TestLastInArray(t *testing.T) {
	t.Log("Testing last element in array")
	array := []int{1, 2, 3, 4, 5, 6, 7}
	result := BSearch(array, 7)
	if result != 6 {
		t.Error("Incorrect result")
	}
}
