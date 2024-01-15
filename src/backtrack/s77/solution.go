package s77

var (
	path []int
	res  [][]int
)

func Combine(n int, k int) [][]int {
	//  初始化全局path,res变量
	path, res = make([]int, 0, k), make([][]int, 0)
	dfs(n, k, 1)
	return res
}

func dfs(n int, k int, start int) {
	if len(path) == k { // 说明已经满足了k个数的要求
		tmp := make([]int, k)
		copy(tmp, path)
		res = append(res, tmp)
		return
	}
	for i := start; i <= n; i++ { // 从start开始，不往回走，避免出现重复组合
		//  减枝，当k-len大于n-i+1,当剩余数加上已经累加的path（i到n,少于等于k的时候遍历就没有必要了）
		//  4-2+1>2
		//  4-3+1=2 ，等于3最后一个
		//  4-4+1<2
		//  i=4，从4开始遍历是不可以的
		if n-i+1+len(path) < k {
			break
		}
		//  从i开始的字符串(i从1开始)
		path = append(path, i)
		dfs(n, k, i+1)
		path = path[:len(path)-1]
	}
}
