package main

import "testing"

func TestTwoSum(t *testing.T) {
	arr, target, want := []int{1, 3, 2, 4, 7, 2, 7, 6, 9, 10}, 10, []int{1, 4}
	if result := TwoSum(arr, target); result[0] != 1 || result[1] != 4 {
		t.Errorf("Want: index{ %d, %d } and get: index{ %d, %d }", want[0], want[1], result[0], result[1])
	}
}

func TestTwoSumVersion1(t *testing.T) {
	arr, target, want := []int{0, 3, 2, 5, 7, 2, 7, 6, 9, 9}, 10, []int{1, 4}
	if result := TwoSumVersion1(arr, target); result[0] != want[0] || result[1] != want[1] {
		t.Errorf("Want: index{ %d, %d } and get: index{ %d, %d }", want[0], want[1], result[0], result[1])
	}
	arr, target, want = []int{4, 4, 11, 13}, 8, []int{0, 1}
	if result := TwoSumVersion1(arr, target); result[0] != want[0] || result[1] != want[1] {
		t.Errorf("Want: index{ %d, %d } and get: index{ %d, %d }", want[0], want[1], result[0], result[1])
	}
}

func TestTwoSumVersion2(t *testing.T) {
	arr, target, want := []int{0, 3, 2, 5, 7, 2, 7, 6, 9, 9}, 10, []int{1, 4}
	if result := TwoSumVersion2(arr, target); result[0] != want[0] || result[1] != want[1] {
		t.Errorf("Want: index{ %d, %d } and get: index{ %d, %d }", want[0], want[1], result[0], result[1])
	}
	arr, target, want = []int{4, 4, 11, 13}, 8, []int{0, 1}
	if result := TwoSumVersion2(arr, target); result[0] != want[0] || result[1] != want[1] {
		t.Errorf("Want: index{ %d, %d } and get: index{ %d, %d }", want[0], want[1], result[0], result[1])
	}
}
