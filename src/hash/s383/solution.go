package s383

// CanConstruct
// @Description 寻找相同字母
// @Author lqc 2024-01-12 10:45:59
// @Param ransomNote
// @Param magazine
// @Return bool
func CanConstruct(ransomNote string, magazine string) bool {
	record := make([]int, 26)
	for _, v := range magazine { // 通过record数据记录 magazine里各个字符出现次数
		record[v-'a']++
	}
	for _, v := range ransomNote { // 遍历ransomNote，在record里对应的字符个数做--操作
		record[v-'a']--
		if record[v-'a'] < 0 { // 如果小于零说明ransomNote里出现的字符，magazine没有
			return false
		}
	}
	return true
}
