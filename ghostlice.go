// Copyright (c) 2014 Aubrey Barnard.  This is free software.  See
// LICENSE for details.

package ghostlice

// MinUint8s finds the minimum of a slice of uint8s and its index
func MinUint8s(nums ...uint8) (min uint8, index int) {
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

// MaxUint8s finds the maximum of a slice of uint8s and its index
func MaxUint8s(nums ...uint8) (max uint8, index int) {
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

// MinMaxUint8s finds both extremes of a slice of uint8s and their
// indices
func MinMaxUint8s(nums ...uint8) (min, max uint8, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumUint8s computes the sum of a slice of uint8s.  The sum must fit
// into a uint8 or it overflows.
func SumUint8s(nums ...uint8) (sum uint8) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillUint8s fills a slice of uint8s with a given value
func FillUint8s(nums []uint8, value uint8) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualUint8s compares two slices of uint8s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualUint8s(nums1, nums2 []uint8) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindUint8 returns the first index of the given value in the slice of
// uint8s
func FindUint8(value uint8, nums ...uint8) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromUint8s makes a slice of float64s by converting a slice of
// uint8s
func Float64sFromUint8s(nums ...uint8) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinUint16s finds the minimum of a slice of uint16s and its index
func MinUint16s(nums ...uint16) (min uint16, index int) {
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

// MaxUint16s finds the maximum of a slice of uint16s and its index
func MaxUint16s(nums ...uint16) (max uint16, index int) {
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

// MinMaxUint16s finds both extremes of a slice of uint16s and their
// indices
func MinMaxUint16s(nums ...uint16) (min, max uint16, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumUint16s computes the sum of a slice of uint16s.  The sum must fit
// into a uint16 or it overflows.
func SumUint16s(nums ...uint16) (sum uint16) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillUint16s fills a slice of uint16s with a given value
func FillUint16s(nums []uint16, value uint16) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualUint16s compares two slices of uint16s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualUint16s(nums1, nums2 []uint16) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindUint16 returns the first index of the given value in the slice of
// uint16s
func FindUint16(value uint16, nums ...uint16) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromUint16s makes a slice of float64s by converting a slice of
// uint16s
func Float64sFromUint16s(nums ...uint16) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinUint32s finds the minimum of a slice of uint32s and its index
func MinUint32s(nums ...uint32) (min uint32, index int) {
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

// MaxUint32s finds the maximum of a slice of uint32s and its index
func MaxUint32s(nums ...uint32) (max uint32, index int) {
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

// MinMaxUint32s finds both extremes of a slice of uint32s and their
// indices
func MinMaxUint32s(nums ...uint32) (min, max uint32, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumUint32s computes the sum of a slice of uint32s.  The sum must fit
// into a uint32 or it overflows.
func SumUint32s(nums ...uint32) (sum uint32) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillUint32s fills a slice of uint32s with a given value
func FillUint32s(nums []uint32, value uint32) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualUint32s compares two slices of uint32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualUint32s(nums1, nums2 []uint32) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindUint32 returns the first index of the given value in the slice of
// uint32s
func FindUint32(value uint32, nums ...uint32) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromUint32s makes a slice of float64s by converting a slice of
// uint32s
func Float64sFromUint32s(nums ...uint32) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinUint64s finds the minimum of a slice of uint64s and its index
func MinUint64s(nums ...uint64) (min uint64, index int) {
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

// MaxUint64s finds the maximum of a slice of uint64s and its index
func MaxUint64s(nums ...uint64) (max uint64, index int) {
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

// MinMaxUint64s finds both extremes of a slice of uint64s and their
// indices
func MinMaxUint64s(nums ...uint64) (min, max uint64, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumUint64s computes the sum of a slice of uint64s.  The sum must fit
// into a uint64 or it overflows.
func SumUint64s(nums ...uint64) (sum uint64) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillUint64s fills a slice of uint64s with a given value
func FillUint64s(nums []uint64, value uint64) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualUint64s compares two slices of uint64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualUint64s(nums1, nums2 []uint64) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindUint64 returns the first index of the given value in the slice of
// uint64s
func FindUint64(value uint64, nums ...uint64) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromUint64s makes a slice of float64s by converting a slice of
// uint64s
func Float64sFromUint64s(nums ...uint64) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinUints finds the minimum of a slice of uints and its index
func MinUints(nums ...uint) (min uint, index int) {
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

// MaxUints finds the maximum of a slice of uints and its index
func MaxUints(nums ...uint) (max uint, index int) {
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

// MinMaxUints finds both extremes of a slice of uints and their
// indices
func MinMaxUints(nums ...uint) (min, max uint, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumUints computes the sum of a slice of uints.  The sum must fit
// into a uint or it overflows.
func SumUints(nums ...uint) (sum uint) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillUints fills a slice of uints with a given value
func FillUints(nums []uint, value uint) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualUints compares two slices of uints for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualUints(nums1, nums2 []uint) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindUint returns the first index of the given value in the slice of
// uints
func FindUint(value uint, nums ...uint) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromUints makes a slice of float64s by converting a slice of
// uints
func Float64sFromUints(nums ...uint) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinInt8s finds the minimum of a slice of int8s and its index
func MinInt8s(nums ...int8) (min int8, index int) {
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

// MaxInt8s finds the maximum of a slice of int8s and its index
func MaxInt8s(nums ...int8) (max int8, index int) {
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

// MinMaxInt8s finds both extremes of a slice of int8s and their
// indices
func MinMaxInt8s(nums ...int8) (min, max int8, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumInt8s computes the sum of a slice of int8s.  The sum must fit
// into a int8 or it overflows.
func SumInt8s(nums ...int8) (sum int8) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillInt8s fills a slice of int8s with a given value
func FillInt8s(nums []int8, value int8) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualInt8s compares two slices of int8s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualInt8s(nums1, nums2 []int8) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindInt8 returns the first index of the given value in the slice of
// int8s
func FindInt8(value int8, nums ...int8) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromInt8s makes a slice of float64s by converting a slice of
// int8s
func Float64sFromInt8s(nums ...int8) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinInt16s finds the minimum of a slice of int16s and its index
func MinInt16s(nums ...int16) (min int16, index int) {
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

// MaxInt16s finds the maximum of a slice of int16s and its index
func MaxInt16s(nums ...int16) (max int16, index int) {
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

// MinMaxInt16s finds both extremes of a slice of int16s and their
// indices
func MinMaxInt16s(nums ...int16) (min, max int16, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumInt16s computes the sum of a slice of int16s.  The sum must fit
// into a int16 or it overflows.
func SumInt16s(nums ...int16) (sum int16) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillInt16s fills a slice of int16s with a given value
func FillInt16s(nums []int16, value int16) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualInt16s compares two slices of int16s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualInt16s(nums1, nums2 []int16) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindInt16 returns the first index of the given value in the slice of
// int16s
func FindInt16(value int16, nums ...int16) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromInt16s makes a slice of float64s by converting a slice of
// int16s
func Float64sFromInt16s(nums ...int16) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinInt32s finds the minimum of a slice of int32s and its index
func MinInt32s(nums ...int32) (min int32, index int) {
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

// MaxInt32s finds the maximum of a slice of int32s and its index
func MaxInt32s(nums ...int32) (max int32, index int) {
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

// MinMaxInt32s finds both extremes of a slice of int32s and their
// indices
func MinMaxInt32s(nums ...int32) (min, max int32, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumInt32s computes the sum of a slice of int32s.  The sum must fit
// into a int32 or it overflows.
func SumInt32s(nums ...int32) (sum int32) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillInt32s fills a slice of int32s with a given value
func FillInt32s(nums []int32, value int32) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualInt32s compares two slices of int32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualInt32s(nums1, nums2 []int32) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindInt32 returns the first index of the given value in the slice of
// int32s
func FindInt32(value int32, nums ...int32) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromInt32s makes a slice of float64s by converting a slice of
// int32s
func Float64sFromInt32s(nums ...int32) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinInt64s finds the minimum of a slice of int64s and its index
func MinInt64s(nums ...int64) (min int64, index int) {
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

// MaxInt64s finds the maximum of a slice of int64s and its index
func MaxInt64s(nums ...int64) (max int64, index int) {
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

// MinMaxInt64s finds both extremes of a slice of int64s and their
// indices
func MinMaxInt64s(nums ...int64) (min, max int64, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumInt64s computes the sum of a slice of int64s.  The sum must fit
// into a int64 or it overflows.
func SumInt64s(nums ...int64) (sum int64) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillInt64s fills a slice of int64s with a given value
func FillInt64s(nums []int64, value int64) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualInt64s compares two slices of int64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualInt64s(nums1, nums2 []int64) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindInt64 returns the first index of the given value in the slice of
// int64s
func FindInt64(value int64, nums ...int64) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromInt64s makes a slice of float64s by converting a slice of
// int64s
func Float64sFromInt64s(nums ...int64) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinInts finds the minimum of a slice of ints and its index
func MinInts(nums ...int) (min int, index int) {
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

// MaxInts finds the maximum of a slice of ints and its index
func MaxInts(nums ...int) (max int, index int) {
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

// MinMaxInts finds both extremes of a slice of ints and their
// indices
func MinMaxInts(nums ...int) (min, max int, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumInts computes the sum of a slice of ints.  The sum must fit
// into a int or it overflows.
func SumInts(nums ...int) (sum int) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillInts fills a slice of ints with a given value
func FillInts(nums []int, value int) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualInts compares two slices of ints for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualInts(nums1, nums2 []int) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindInt returns the first index of the given value in the slice of
// ints
func FindInt(value int, nums ...int) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromInts makes a slice of float64s by converting a slice of
// ints
func Float64sFromInts(nums ...int) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinFloat32s finds the minimum of a slice of float32s and its index
func MinFloat32s(nums ...float32) (min float32, index int) {
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

// MaxFloat32s finds the maximum of a slice of float32s and its index
func MaxFloat32s(nums ...float32) (max float32, index int) {
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

// MinMaxFloat32s finds both extremes of a slice of float32s and their
// indices
func MinMaxFloat32s(nums ...float32) (min, max float32, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumFloat32s computes the sum of a slice of float32s.  The sum must fit
// into a float32 or it overflows.
func SumFloat32s(nums ...float32) (sum float32) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillFloat32s fills a slice of float32s with a given value
func FillFloat32s(nums []float32, value float32) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualFloat32s compares two slices of float32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualFloat32s(nums1, nums2 []float32) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindFloat32 returns the first index of the given value in the slice of
// float32s
func FindFloat32(value float32, nums ...float32) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromFloat32s makes a slice of float64s by converting a slice of
// float32s
func Float64sFromFloat32s(nums ...float32) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

// MinFloat64s finds the minimum of a slice of float64s and its index
func MinFloat64s(nums ...float64) (min float64, index int) {
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

// MaxFloat64s finds the maximum of a slice of float64s and its index
func MaxFloat64s(nums ...float64) (max float64, index int) {
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

// MinMaxFloat64s finds both extremes of a slice of float64s and their
// indices
func MinMaxFloat64s(nums ...float64) (min, max float64, minIndex, maxIndex int) {
	if len(nums) == 0 {
		return
	}
	min = nums[0]
	max = nums[0]
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

// SumFloat64s computes the sum of a slice of float64s.  The sum must fit
// into a float64 or it overflows.
func SumFloat64s(nums ...float64) (sum float64) {
	if len(nums) == 0 {
		return
	}
	for _, value := range nums {
		sum += value
	}
	return
}

// FillFloat64s fills a slice of float64s with a given value
func FillFloat64s(nums []float64, value float64) {
	for i := range nums {
		nums[i] = value
	}
}

// EqualFloat64s compares two slices of float64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ
func EqualFloat64s(nums1, nums2 []float64) (equal bool, index int) {
	length := len(nums1)
	if length != len(nums2) {
		return false, -1
	}
	for i := 0; i < length; i++ {
		if nums1[i] != nums2[i] {
			return false, i
		}
	}
	return true, length
}

// FindFloat64 returns the first index of the given value in the slice of
// float64s
func FindFloat64(value float64, nums ...float64) (index int, found bool) {
	for i, v := range nums {
		if v == value {
			return i, true
		}
	}
	return
}

// Float64sFromFloat64s makes a slice of float64s by converting a slice of
// float64s
func Float64sFromFloat64s(nums ...float64) (floats []float64) {
	length := len(nums)
	floats = make([]float64, length)
	for i := 0; i < length; i++ {
		floats[i] = float64(nums[i])
	}
	return
}

