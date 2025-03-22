package search

func BinarySearch(nums []int, target int) int {
	return binarySearchHelper(nums, target, 0, len(nums)-1)
}

func binarySearchHelper(nums []int, target, low, high int) int {
	// Check the base case
	if low > high {
		return -1
	}

	// Find the midpoint of a given array
	mid := low + (high-low)/2

	// CASE01: Check if the midpoint is the target
	// CASE02: Check if midpoint is less than target if true then decrement the midpoint
	// CASE03: Check if midpoint is greater than target if true then increment the midpoint
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearchHelper(nums, target, low, mid-1)
	} else {
		return binarySearchHelper(nums, target, mid+1, high)
	}
}
