// Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
// LICENSE for details.

// Tests ghostlice.go with at least one test per templated function.

package ghostlice

import (
	. "testing"
)

var emptyInts = []int{}

var ints = []int{
	2, -64, 93, -38, 33, 82, 67, 40, 82, 42,
	38, 53, 27, 46, 49, 82, -98, 38, 46, 6,
	-84, -21, -93, 53, 86, 74, 2, 18, -76, 90,
	-90, -4, -67, 24, 38, -5, 84,
}

func Test_MinInts(t *T) {
	testData := [][]int{emptyInts, ints[:1], ints}
	expectedMins := []int{0, 2, -98}
	expectedIndexes := []int{0, 0, 16}
	for test := 0; test < len(testData); test++ {
		min, index := MinInts(testData[test]...)
		if min != expectedMins[test] {
			t.Errorf("Wrong minimum: expected %v != %v", expectedMins[test], min)
		}
		if index != expectedIndexes[test] {
			t.Errorf("Wrong index of minimum: expected %v != %v", expectedIndexes[test], index)
		}
	}
}

func Test_MaxInts(t *T) {
	testData := [][]int{emptyInts, ints[:1], ints}
	expectedMaxs := []int{0, 2, 93}
	expectedIndexes := []int{0, 0, 2}
	for test := 0; test < len(testData); test++ {
		max, index := MaxInts(testData[test]...)
		if max != expectedMaxs[test] {
			t.Errorf("Wrong maximum: expected %v != %v", expectedMaxs[test], max)
		}
		if index != expectedIndexes[test] {
			t.Errorf("Wrong index of maximum: expected %v != %v", expectedIndexes[test], index)
		}
	}
}

func Test_MinMaxInts(t *T) {
	testData := [][]int{emptyInts, ints[:1], ints}
	expectedMins := []int{0, 2, -98}
	expectedMaxs := []int{0, 2, 93}
	expectedMinIndexes := []int{0, 0, 16}
	expectedMaxIndexes := []int{0, 0, 2}
	for test := 0; test < len(testData); test++ {
		min, max, minIndex, maxIndex := MinMaxInts(testData[test]...)
		if min != expectedMins[test] {
			t.Errorf("Wrong minimum: expected %v != %v", expectedMins[test], min)
		}
		if max != expectedMaxs[test] {
			t.Errorf("Wrong maximum: expected %v != %v", expectedMaxs[test], max)
		}
		if minIndex != expectedMinIndexes[test] {
			t.Errorf("Wrong index of minimum: expected %v != %v", expectedMinIndexes[test], minIndex)
		}
		if maxIndex != expectedMaxIndexes[test] {
			t.Errorf("Wrong index of maximum: expected %v != %v", expectedMaxIndexes[test], maxIndex)
		}
	}
}

func Test_SumInts(t *T) {
	testData := [][]int{emptyInts, ints[:1], ints}
	expectedSums := []int{0, 2, 655}
	for test := 0; test < len(testData); test++ {
		sum := SumInts(testData[test]...)
		if sum != expectedSums[test] {
			t.Errorf("Wrong sum: expected %v != %v", expectedSums[test], sum)
		}
	}
}

func Test_FillInts(t *T) {
	sizes := []int{len(emptyInts), len(ints)}
	fills := []int{0, -48}
	for test := 0; test < len(sizes); test++ {
		nums := make([]int, sizes[test])
		FillInts(nums, fills[test])
		min, max, _, _ := MinMaxInts(nums...)
		if min != fills[test] || max != fills[test] {
			t.Errorf("Wrong fill: expected %v != (%v or %v)", fills[test], min, max)
		}
	}
}

func Test_EqualInts(t *T) {
	// Make a slice with the same contents except the last
	intsCopy := make([]int, len(ints))
	copy(intsCopy, ints)
	intsCopy[len(intsCopy)-1] = 88
	// Test cases.  False cases: empty vs. non, len diff, first element
	// diff, middle element diff, last element diff
	ints1 := [][]int{emptyInts, ints[:1], ints, emptyInts, ints[:5], ints[2:7], ints[5:9], ints}
	ints2 := [][]int{emptyInts, ints[:1], ints, ints[:1], ints[:7], ints[3:8], ints[8:12], intsCopy}
	equals := []bool{true, true, true, false, false, false, false, false}
	indices := []int{0, 1, 37, -1, -1, 0, 1, 36}
	for test := 0; test < len(equals); test++ {
		equality, index := EqualInts(ints1[test], ints2[test])
		if equality != equals[test] {
			t.Errorf("Wrong equality in test %v: expected %v != %v", test, equals[test], equality)
		}
		if index != indices[test] {
			t.Errorf("Wrong index in test %v: expected %v != %v", test, indices[test], index)
		}
	}
}

func Test_FindInt(t *T) {
	testData := [][]int{emptyInts, ints[:1], ints[:1], ints, ints, ints}
	targets := []int{363, 2, -880, -67, 38, 1}
	expectedIndexes := []int{0, 0, 0, 32, 10, 0}
	expectedFounds := []bool{false, true, false, true, true, false}
	for test := 0; test < len(testData); test++ {
		index, found := FindInt(targets[test], testData[test]...)
		if index != expectedIndexes[test] {
			t.Errorf("Wrong index in test %v: expected %v != %v", targets[test], expectedIndexes[test], index)
		}
		if found != expectedFounds[test] {
			t.Errorf("Wrong found in test %v: expected %v != %v", targets[test], expectedFounds[test], found)
		}
	}
}
