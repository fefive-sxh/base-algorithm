package find_dup_num

import "testing"

func createFindNoMoveCase(t *testing.T, d *dupNumSlice) {
	t.Helper()
	if ans := FindDupNoMove(d.nums); ans != d.expect {
		t.Fatalf("%v expected %d, but got %d", d.nums, d.expect, ans)
	}
}

func TestFindDupNoMove(t *testing.T) {
	createFindNoMoveCase(t, &dupNumSlice{[]int{1, 2, 3, 5, 4, 1}, 1})
	createFindNoMoveCase(t, &dupNumSlice{[]int{1, 2, 3, 2, 3, 5}, 2}) // wrong case
}

func BenchmarkFindDupNoMove(b *testing.B) {
	nums := []int{1, 3, 2, 4, 6, 5, 7, 9, 8, 7}
	// 重置计时器
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindDupNoMove(nums)
	}
}

/*
BenchmarkFindDupNoMove-4        17787496                63.8 ns/op             0 B/op          0 allocs/op
*/
