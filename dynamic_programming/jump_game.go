package dynamic_programming

/*
A common interview question asks: "How many distinct ways can you climb a staircase with n steps, given that you can
take either 1 or 2 steps at a time?" This problem can be solved with both recursion and dynamic programming.
*/

func JumpBottomUp(n int) int {
	jumps := map[int]int{
		0: 1,
		1: 1,
	}

	for i := 2; i <= n; i++ {
		jumps[i] = jumps[i-1] + jumps[i-2]
	}

	return jumps[n]
}

func JumpRecursive(n int) int {
	if n <= 1 {
		return 1
	}

	return JumpRecursive(n-1) + JumpRecursive(n-2)
}

func JumpRecursiveMemo(n int, memo map[int]int) int {
	if n <= 1 {
		memo[n] = 1
		return memo[n]
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = JumpRecursive(n-1) + JumpRecursive(n-2)
	return memo[n]
}
