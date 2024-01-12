package s59

func GenerateMatrix(n int) [][]int {
	startX, startY := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startX, startY
		//行数不变 列数在变
		for j = startY; j < n-offset; j++ {
			res[startX][j] = count
			count++
		}
		//列数不变是j 行数变
		for i = startX; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//行数不变 i 列数变 j--
		for ; j > startY; j-- {
			res[i][j] = count
			count++
		}
		//列不变 行变
		for ; i > startX; i-- {
			res[i][j] = count
			count++
		}
		startX++
		startY++
		offset++
		loop--
	}

	// 如果n为奇数的话，需要单独给矩阵最中间的位置赋值
	if n%2 == 1 {
		res[center][center] = n * n
	}
	return res
}
