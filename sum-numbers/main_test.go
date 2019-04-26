package main

import "testing"

func assertNumbers(
	t *testing.T,
	numbers []int,
	k int,
	expectedResult bool,
) {
	result := IsAddUpTo(numbers, k)

	if result && !expectedResult {
		t.Errorf("Numbers %v should not add up to %d", numbers, k)
	}
	if !result && expectedResult {
		t.Errorf("Numbers %v should add up to %d", numbers, k)
	}
}

func TestEmptyList(t *testing.T) {
	assertNumbers(t, []int{}, 1, false)
}

func TestListWithOneNotKElement(t *testing.T) {
	assertNumbers(t, []int{1}, 5, false)
}

func TestListWithOneWithKElement(t *testing.T) {
	assertNumbers(t, []int{5}, 5, true)
}

func TestListWithTwoElements(t *testing.T) {
	assertNumbers(t, []int{1, 2}, 5, false)
}

func TestListWithTwoElementsAddsUp(t *testing.T) {
	assertNumbers(t, []int{1, 4}, 5, true)
}

func TestListWithThreeElementsAddsUp(t *testing.T) {
	assertNumbers(t, []int{1, 2, 4}, 5, true)
}

func TestListWithFourElementsAddsUp(t *testing.T) {
	assertNumbers(t, []int{10, 15, 3, 7}, 17, true)
}
