package find_dup_num

// 题目2: 不修改数组找出重复的数字

// 在一个长度为 n+1 的数组里的所有数字都在 1~n 的范围内,
// 所以至少有一个数字是重复的, 请找出任意重复的数字
// 但不能修改输入的数组

// 思路:
// 若没有重复, 则 1 ~ n+1 仅出现一次, 由于重复了, 可以考虑二分 m := (1 + n ) / 2
// 判断 小于 m 的个数 是否 大于m个, 若大于m 则在左半数组, 否则 在右半数组, 依此类推 直到剩下一个
func FindDupNoMove(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		count := countRange(nums, start, mid)

		// 结束条件
		if start == end {
			if count > 1 {
				return start
			} else {
				break
			}
		}
		// 判断左边还是右边
		if count > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return -1
}

//  统计切片中 数字在 [start, end] 中的个数
func countRange(nums []int, start, end int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}
