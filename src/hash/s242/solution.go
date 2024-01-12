package s242

// IsAnagram1
// @Description 记录每个字母出现几次，直接比对就ok
// @Author lqc 2024-01-11 19:24:36
// @Param s
// @Param t
// @Return bool
func IsAnagram1(s string, t string) bool {
	record := [26]int{}

	for _, r := range s {
		record[r-rune('a')]++
	}
	for _, r := range t {
		record[r-rune('a')]--
	}

	return record == [26]int{} // record数组如果有的元素不为零0，说明字符串s和t 一定是谁多了字符或者谁少了字符。
}

func IsAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int{}
	//  1.遍历s
	for index := 0; index < len(s); index++ {
		if s[index] == t[index] {
			continue
		}
		// 2.每种字母出现的次数
		sCharIndex := s[index] - 'a'
		records[sCharIndex]++
		tCharIndex := t[index] - 'a'
		records[tCharIndex]--
	}
	for _, record := range records {
		if record != 0 {
			return false
		}
	}
	return true
}
