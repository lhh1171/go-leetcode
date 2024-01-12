package s202

// IsHappy
// @Description 观察是否存在相同的数
// @Author lqc 2024-01-11 19:46:39
// @Param n
// @Return bool
func IsHappy(n int) bool {
	m := make(map[int]bool)
	//  求和等于一，或者该求和存在过
	for n != 1 && !m[n] {
		n, m[n] = getSum(n), true
	}
	return n == 1
}

// getSum
// @Description 给该数字每一个位置的数字平方求和
// @Author lqc 2024-01-11 19:48:18
// @Param n
// @Return int
func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}
