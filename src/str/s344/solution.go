package s344

// ReverseString
// @Description 双指针法反转字符串
// @Author lqc 2024-01-12 10:54:16
// @Param s
func ReverseString(s []byte) string {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}
