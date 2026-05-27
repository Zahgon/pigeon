// Copied from https://github.com/stretchr/testify

// Copyright (c) 2012-2020 Mat Ryer, Tyler Bunnell and contributors. All rights reserved.
// Use of this source code is governed by an MIT-style license that can be found in
// the THIRD-PARTY-NOTICES file.

package testutils

// IsEmpty gets whether the specified object is considered empty or not.
func IsEmpty(object interface{}) bool {
	_ = "STUB: not implemented"
	// get nil case out of the way
	return false
}

// collection types are empty when they have no element

// pointers are empty if nil or if the value they point to is empty

// for all other types, compare against the zero value
// array types are empty when they match their zero-initialized state

// IsList checks that the provided value is array or slice.
func IsList(list interface{}) (ok bool) { _ = "STUB: not implemented"; return false }

// DiffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B.
// If some element is present multiple times, each instance is counted separately (e.g. if something is 2x in A and
// 5x in B, it will be 0x in extraA and 3x in extraB). The order of items in both lists is ignored.
func DiffLists(listA, listB interface{}) (extraA, extraB []interface{}) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mark indexes in bValue that we already used

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool { _ = "STUB: not implemented"; return false }

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// ElementsMatch([1, 3, 2, 3], [1, 3, 3, 2]).
func ElementsMatch(listA interface{}, listB interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

func isFunction(arg interface{}) bool { _ = "STUB: not implemented"; return false }

// ValidateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func ValidateEqualArgs(expected, actual interface{}) error { _ = "STUB: not implemented"; return nil }

// Equal asserts that two objects are equal.
//
//	Equal(123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	_ = "STUB: not implemented"
	return false
}
