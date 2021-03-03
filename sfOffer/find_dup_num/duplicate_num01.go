package find_dup_num

import (
	"math/rand"
	"time"
)

// 在一个长度为 n 的数组中, 值都在 0 ~ n-1 范围内,
// 数组中的某些数字是重复的, 但不知道有几个重复, 以及重复了几次
// 请找出数组中任意一个重复的数字

// eg. [2,3,1,0,2,3,5], 输出为: 2 或者 3


// 思路: 由于 数组长度为n, 而且数组中的值 也在 [0, n), 故而可以使 值和下标索引归位,
// 当遍历到某一个值的时候, 若其值 和 数组索引中的值相等时, 则返回此值

// 1. 首先构造一个随机的切片, 切片中值在[0, len(nums)-1] 之间随机
func GenerateRandomArray(n int) []int {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		// n-1 保证 至少有一个 是重复的
		nums = append(nums, rand.Intn(n-1))
	}
	return nums
}


func FindDuplication(nums []int) int {
	if len(nums) == 0 {
		return -1
		// -1 代表错误
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		for n != i {
			if nums[n] == n {
				return n
			}else {
				nums[i], nums[n] = nums[n], nums[i]
				n = nums[i]
			}
		}
	}
	return -1
}

// 若不修改 原来数组
// 1. 则可以通过 创建一个 map 来记录出现过的数字
func FindDupNumByBoolMap(nums []int) int {
	m := make(map[int]bool, len(nums))

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}else {
			m[nums[i]] = true
		}
	}
	return -1
}

// 测试发现 使用第一种map 会随着数组长度增加, 而增加内存分配和操作时长
// 2. 使用 空结构体map
func FindDupNumByStructMap(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		}
		m[nums[i]] = struct{}{}
	}
	return -1
}

