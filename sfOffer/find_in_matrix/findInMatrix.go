package find_in_matrix

// 在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
// 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func FindInMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}

	row, col := 0, len(matrix[0])-1
	for {
		if row >= len(matrix) || col < 0 {
			break
		}

		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
