// Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
// LICENSE for details.

// This file is a template and not meant to be used directly.

// MinATypes finds the minimum of a slice of ATypes and its index
func MinATypes(nums ...AType) (min AType, index int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			index = i
		}
	}
	return
}

// MaxATypes finds the maximum of a slice of ATypes and its index
func MaxATypes(nums ...AType) (max AType, index int) {
	if len(nums) == 0 {
		return
	}
	max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			index = i
		}
	}
	return
}

// MinMaxATypes finds both extremes of a slice of ATypes and their
// indices
func MinMaxATypes(nums ...AType) (min, max AType, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
			minIndex = i
		}
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}
	return
}

// SumATypes computes the sum of a slice of ATypes.  The sum must fit
// into a AType or it overflows.
func SumATypes(nums ...AType) (sum AType) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillATypes fills a slice of ATypes with a given value
func FillATypes(nums []AType, value AType) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualATypes compares two slices of ATypes for equality of contents
func EqualATypes(nums1, nums2 []AType) (equal bool) {
	length := len(nums1)
	if length != len(nums2) {
		return
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return
		}
	}
	return true
}

// FindAType returns the first index of the given value in the slice of
// ATypes
func FindAType(value AType, nums ...AType) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

