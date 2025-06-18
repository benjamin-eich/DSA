package backtrack

/*
Generating All Permutations of a List
This classic problem uses backtracking to generate every possible ordering of a list.
*/

func Permutations(nums []int) [][]int {
	result := make([][]int, 0)
	var backtrack func(start int)

	backtrack = func(start int) {
		if start == len(nums) {
			// Make a copy of nums and add it to the result
			perm := make([]int, len(nums))
			copy(perm, nums)
			result = append(result, perm)
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start] // Swap
			backtrack(start + 1)
			nums[start], nums[i] = nums[i], nums[start] // Swap back
		}
	}

	backtrack(0)
	return result
}
