package s216

var (
	res  [][]int
	path []int
)

func CombinationSum(k int, n int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, k)
	dfs(k, n, 1, 0)
	return res
}

func dfs(k, n int, start int, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}
	for i := start; i <= 9; i++ {
		if sum+i > n || 9-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		dfs(k, n, i+1, sum+i)
		path = path[:len(path)-1]
	}
}
