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
func FillUint8s(nums []uint8, value uint8) []uint8 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualUint8s compares two slices of uint8s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualUint8s(nums1, nums2 []uint8) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindUint8(value uint8, nums ...uint8) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddUint8sSc adds the given number (scalar) to each element
func AddUint8sSc(addend uint8, nums ...uint8) (sums []uint8) {
	length := len(nums)
	sums = make([]uint8, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubUint8sSc subtracts the given number (scalar) from each element
func SubUint8sSc(subtrahend uint8, nums ...uint8) (differences []uint8) {
	length := len(nums)
	differences = make([]uint8, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulUint8sSc multiplies each element by the given number (scalar)
func MulUint8sSc(multiplier uint8, nums ...uint8) (products []uint8) {
	length := len(nums)
	products = make([]uint8, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivUint8sSc divides each element by the given number (scalar)
func DivUint8sSc(divisor uint8, nums ...uint8) (quotients []uint8) {
	length := len(nums)
	quotients = make([]uint8, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillUint16s(nums []uint16, value uint16) []uint16 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualUint16s compares two slices of uint16s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualUint16s(nums1, nums2 []uint16) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindUint16(value uint16, nums ...uint16) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddUint16sSc adds the given number (scalar) to each element
func AddUint16sSc(addend uint16, nums ...uint16) (sums []uint16) {
	length := len(nums)
	sums = make([]uint16, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubUint16sSc subtracts the given number (scalar) from each element
func SubUint16sSc(subtrahend uint16, nums ...uint16) (differences []uint16) {
	length := len(nums)
	differences = make([]uint16, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulUint16sSc multiplies each element by the given number (scalar)
func MulUint16sSc(multiplier uint16, nums ...uint16) (products []uint16) {
	length := len(nums)
	products = make([]uint16, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivUint16sSc divides each element by the given number (scalar)
func DivUint16sSc(divisor uint16, nums ...uint16) (quotients []uint16) {
	length := len(nums)
	quotients = make([]uint16, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillUint32s(nums []uint32, value uint32) []uint32 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualUint32s compares two slices of uint32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualUint32s(nums1, nums2 []uint32) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindUint32(value uint32, nums ...uint32) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddUint32sSc adds the given number (scalar) to each element
func AddUint32sSc(addend uint32, nums ...uint32) (sums []uint32) {
	length := len(nums)
	sums = make([]uint32, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubUint32sSc subtracts the given number (scalar) from each element
func SubUint32sSc(subtrahend uint32, nums ...uint32) (differences []uint32) {
	length := len(nums)
	differences = make([]uint32, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulUint32sSc multiplies each element by the given number (scalar)
func MulUint32sSc(multiplier uint32, nums ...uint32) (products []uint32) {
	length := len(nums)
	products = make([]uint32, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivUint32sSc divides each element by the given number (scalar)
func DivUint32sSc(divisor uint32, nums ...uint32) (quotients []uint32) {
	length := len(nums)
	quotients = make([]uint32, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillUint64s(nums []uint64, value uint64) []uint64 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualUint64s compares two slices of uint64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualUint64s(nums1, nums2 []uint64) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindUint64(value uint64, nums ...uint64) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddUint64sSc adds the given number (scalar) to each element
func AddUint64sSc(addend uint64, nums ...uint64) (sums []uint64) {
	length := len(nums)
	sums = make([]uint64, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubUint64sSc subtracts the given number (scalar) from each element
func SubUint64sSc(subtrahend uint64, nums ...uint64) (differences []uint64) {
	length := len(nums)
	differences = make([]uint64, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulUint64sSc multiplies each element by the given number (scalar)
func MulUint64sSc(multiplier uint64, nums ...uint64) (products []uint64) {
	length := len(nums)
	products = make([]uint64, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivUint64sSc divides each element by the given number (scalar)
func DivUint64sSc(divisor uint64, nums ...uint64) (quotients []uint64) {
	length := len(nums)
	quotients = make([]uint64, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillUints(nums []uint, value uint) []uint {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualUints compares two slices of uints for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualUints(nums1, nums2 []uint) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindUint(value uint, nums ...uint) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddUintsSc adds the given number (scalar) to each element
func AddUintsSc(addend uint, nums ...uint) (sums []uint) {
	length := len(nums)
	sums = make([]uint, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubUintsSc subtracts the given number (scalar) from each element
func SubUintsSc(subtrahend uint, nums ...uint) (differences []uint) {
	length := len(nums)
	differences = make([]uint, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulUintsSc multiplies each element by the given number (scalar)
func MulUintsSc(multiplier uint, nums ...uint) (products []uint) {
	length := len(nums)
	products = make([]uint, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivUintsSc divides each element by the given number (scalar)
func DivUintsSc(divisor uint, nums ...uint) (quotients []uint) {
	length := len(nums)
	quotients = make([]uint, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillInt8s(nums []int8, value int8) []int8 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualInt8s compares two slices of int8s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualInt8s(nums1, nums2 []int8) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindInt8(value int8, nums ...int8) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddInt8sSc adds the given number (scalar) to each element
func AddInt8sSc(addend int8, nums ...int8) (sums []int8) {
	length := len(nums)
	sums = make([]int8, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubInt8sSc subtracts the given number (scalar) from each element
func SubInt8sSc(subtrahend int8, nums ...int8) (differences []int8) {
	length := len(nums)
	differences = make([]int8, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulInt8sSc multiplies each element by the given number (scalar)
func MulInt8sSc(multiplier int8, nums ...int8) (products []int8) {
	length := len(nums)
	products = make([]int8, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivInt8sSc divides each element by the given number (scalar)
func DivInt8sSc(divisor int8, nums ...int8) (quotients []int8) {
	length := len(nums)
	quotients = make([]int8, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillInt16s(nums []int16, value int16) []int16 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualInt16s compares two slices of int16s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualInt16s(nums1, nums2 []int16) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindInt16(value int16, nums ...int16) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddInt16sSc adds the given number (scalar) to each element
func AddInt16sSc(addend int16, nums ...int16) (sums []int16) {
	length := len(nums)
	sums = make([]int16, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubInt16sSc subtracts the given number (scalar) from each element
func SubInt16sSc(subtrahend int16, nums ...int16) (differences []int16) {
	length := len(nums)
	differences = make([]int16, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulInt16sSc multiplies each element by the given number (scalar)
func MulInt16sSc(multiplier int16, nums ...int16) (products []int16) {
	length := len(nums)
	products = make([]int16, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivInt16sSc divides each element by the given number (scalar)
func DivInt16sSc(divisor int16, nums ...int16) (quotients []int16) {
	length := len(nums)
	quotients = make([]int16, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillInt32s(nums []int32, value int32) []int32 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualInt32s compares two slices of int32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualInt32s(nums1, nums2 []int32) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindInt32(value int32, nums ...int32) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddInt32sSc adds the given number (scalar) to each element
func AddInt32sSc(addend int32, nums ...int32) (sums []int32) {
	length := len(nums)
	sums = make([]int32, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubInt32sSc subtracts the given number (scalar) from each element
func SubInt32sSc(subtrahend int32, nums ...int32) (differences []int32) {
	length := len(nums)
	differences = make([]int32, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulInt32sSc multiplies each element by the given number (scalar)
func MulInt32sSc(multiplier int32, nums ...int32) (products []int32) {
	length := len(nums)
	products = make([]int32, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivInt32sSc divides each element by the given number (scalar)
func DivInt32sSc(divisor int32, nums ...int32) (quotients []int32) {
	length := len(nums)
	quotients = make([]int32, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillInt64s(nums []int64, value int64) []int64 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualInt64s compares two slices of int64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualInt64s(nums1, nums2 []int64) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindInt64(value int64, nums ...int64) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddInt64sSc adds the given number (scalar) to each element
func AddInt64sSc(addend int64, nums ...int64) (sums []int64) {
	length := len(nums)
	sums = make([]int64, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubInt64sSc subtracts the given number (scalar) from each element
func SubInt64sSc(subtrahend int64, nums ...int64) (differences []int64) {
	length := len(nums)
	differences = make([]int64, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulInt64sSc multiplies each element by the given number (scalar)
func MulInt64sSc(multiplier int64, nums ...int64) (products []int64) {
	length := len(nums)
	products = make([]int64, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivInt64sSc divides each element by the given number (scalar)
func DivInt64sSc(divisor int64, nums ...int64) (quotients []int64) {
	length := len(nums)
	quotients = make([]int64, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillInts(nums []int, value int) []int {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualInts compares two slices of ints for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualInts(nums1, nums2 []int) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindInt(value int, nums ...int) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddIntsSc adds the given number (scalar) to each element
func AddIntsSc(addend int, nums ...int) (sums []int) {
	length := len(nums)
	sums = make([]int, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubIntsSc subtracts the given number (scalar) from each element
func SubIntsSc(subtrahend int, nums ...int) (differences []int) {
	length := len(nums)
	differences = make([]int, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulIntsSc multiplies each element by the given number (scalar)
func MulIntsSc(multiplier int, nums ...int) (products []int) {
	length := len(nums)
	products = make([]int, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivIntsSc divides each element by the given number (scalar)
func DivIntsSc(divisor int, nums ...int) (quotients []int) {
	length := len(nums)
	quotients = make([]int, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillFloat32s(nums []float32, value float32) []float32 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualFloat32s compares two slices of float32s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualFloat32s(nums1, nums2 []float32) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindFloat32(value float32, nums ...float32) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddFloat32sSc adds the given number (scalar) to each element
func AddFloat32sSc(addend float32, nums ...float32) (sums []float32) {
	length := len(nums)
	sums = make([]float32, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubFloat32sSc subtracts the given number (scalar) from each element
func SubFloat32sSc(subtrahend float32, nums ...float32) (differences []float32) {
	length := len(nums)
	differences = make([]float32, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulFloat32sSc multiplies each element by the given number (scalar)
func MulFloat32sSc(multiplier float32, nums ...float32) (products []float32) {
	length := len(nums)
	products = make([]float32, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivFloat32sSc divides each element by the given number (scalar)
func DivFloat32sSc(divisor float32, nums ...float32) (quotients []float32) {
	length := len(nums)
	quotients = make([]float32, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
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
func FillFloat64s(nums []float64, value float64) []float64 {
	for i := range nums {
		nums[i] = value
	}
	return nums
}

// EqualFloat64s compares two slices of float64s for equality of contents
// and returns the index of the first difference or -1 if the lengths
// differ (or -2 if one of the slices is nil)
func EqualFloat64s(nums1, nums2 []float64) (equal bool, index int) {
	// Handle nil
	if nums1 == nil && nums2 == nil {
		return true, 0
	} else if nums1 == nil || nums2 == nil {
		return false, -2
	}
	// Handle non-nil slices
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
func FindFloat64(value float64, nums ...float64) (found bool, index int) {
	for i, v := range nums {
		if v == value {
			return true, i
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

// AddFloat64sSc adds the given number (scalar) to each element
func AddFloat64sSc(addend float64, nums ...float64) (sums []float64) {
	length := len(nums)
	sums = make([]float64, length)
	for i := 0; i < length; i++ {
		sums[i] = nums[i] + addend
	}
	return
}

// SubFloat64sSc subtracts the given number (scalar) from each element
func SubFloat64sSc(subtrahend float64, nums ...float64) (differences []float64) {
	length := len(nums)
	differences = make([]float64, length)
	for i := 0; i < length; i++ {
		differences[i] = nums[i] - subtrahend
	}
	return
}

// MulFloat64sSc multiplies each element by the given number (scalar)
func MulFloat64sSc(multiplier float64, nums ...float64) (products []float64) {
	length := len(nums)
	products = make([]float64, length)
	for i := 0; i < length; i++ {
		products[i] = nums[i] * multiplier
	}
	return

}

// DivFloat64sSc divides each element by the given number (scalar)
func DivFloat64sSc(divisor float64, nums ...float64) (quotients []float64) {
	length := len(nums)
	quotients = make([]float64, length)
	for i := 0; i < length; i++ {
		quotients[i] = nums[i] / divisor
	}
	return
}

