package find_in_matrix

import "testing"

type findMatrix struct {
	matrix     [][]int
	targetNums []int
	isIn       bool
}

func createFindMatrix(t *testing.T, f *findMatrix) {
	t.Helper()
	for _, target := range f.targetNums {
		if ok := FindInMatrix(f.matrix, target); f.isIn != ok {
			// t.Error/t.Errorf 遇到错误不停, 继续执行其他测试用例
			// t.Fatal/t.Fatalf 遇错即停
			t.Errorf("the target: %d, the matrix: %v, expect: %v, got %v", target, f.matrix, f.isIn, ok)
		}
	}
}

func TestFindInMatrix(t *testing.T) {
	createFindMatrix(t, &findMatrix{
		matrix:     [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}},
		targetNums: []int{0, 3, 8, 11}, // 边界条件
		isIn:       true,
	})

	createFindMatrix(t, &findMatrix{
		matrix:     [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}},
		targetNums: []int{-1, 12}, // 错误条件
		isIn:       false,
	})
}
