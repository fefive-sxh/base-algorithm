package find_in_matrix

import "testing"

type findMatrix struct {
	matrix [][]int
	target int
	isIn   bool
}

func createFindMatrix(t *testing.T, f *findMatrix) {
	t.Helper()
	if ok := FindInMatrix(f.matrix, f.target); f.isIn != ok {
		t.Fatalf("the target: %d, the matrix: %v, expect: %v, got %v", f.target, f.matrix, f.isIn, ok)
	}
}

func TestFindInMatrix(t *testing.T) {
	createFindMatrix(t, &findMatrix{
		matrix: [][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}},
		target: 0,
		isIn:   true,
	})

	createFindMatrix(t, &findMatrix{
		matrix: [][]int{{1, 2, 3}, {4, 5, 7}, {8, 10, 12}},
		target: 11,
		isIn:   false,
	})
}
