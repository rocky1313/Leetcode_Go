

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */
// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for key, val := range nums {
		if index, ok := m[target-val]; ok {
			return []int{key, index}
		} else {
			m[val] = key
		}
	}
	return []int{}
}

// @lc code=end
