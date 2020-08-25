package peak_element

/**
 * Solution to https://leetcode.com/problems/find-peak-element/
 * Finds a single peak elemnt in a 1-D array recursively in O(log(n))
 * @returns index of a peak element
 */
func FindPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	if n == 2 {
		if nums[1] > nums[0] {
			return 1
		}
		return 0
	}
	if nums[n/2] < nums[(n/2)-1] {
		return FindPeakElement(nums[0:n/2])
	} else if nums[n/2] < nums[(n/2)+1] {
		return (n/2) + 1 + FindPeakElement(nums[(n/2)+1 : n])
	} else {
		return n/2
	}
}
