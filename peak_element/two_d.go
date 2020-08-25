package peak_element

/**
 * Finds a peak in given 2-D array
 * @returns x,y coordinates of the peak
 */
func FindHillTop(nums [][]int) (int, int) {
	n := len(nums)
	max := findMax(nums[n/2])
	if n == 1 {
		return 0, max
	}
	if n == 2 {
		if nums[1][max] > nums[0][max] {
			return 1, max
		}
		return 0, findMax(nums[0])
	}
	if nums[n/2][max] < nums[(n/2)-1][max] { // left half
		return  FindHillTop(nums[0:n/2])
	} else if nums[n/2][max] < nums[(n/2)+1][max] {
		x, y := FindHillTop(nums[(n/2)+1 : n])
		return (n/2)+1+x, y
	} else {
		return n/2, max
	}
}

/**
 * @returns index of global maxima of 1-D array
 */
func findMax(nums []int) int {
	var index int
	max := nums[0]
	for i, v := range nums {
		if v > max {
			index = i
		}
	}
	return index
}