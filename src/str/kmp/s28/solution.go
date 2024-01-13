package s28

import (
	"fmt"
	"strconv"
)

// 方法一:前缀表使用减1实现

// getNext1 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNext1(next []int, model string) {
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j

	//  遍历模式串，从1开始
	for i := 1; i < len(model); i++ {
		//  当j大于等于0，
		//  遍历的这个位置和
		//  之前匹配的最大公共部分的下一个
		//  进行比较如果相等，
		//  最大相同前后缀又加一
		for j >= 0 && model[i] != model[j+1] {
			j = next[j]
		}
		if model[i] == model[j+1] {
			j++
		}
		fmt.Println("i等于" + strconv.Itoa(i) + ",j等于" + strconv.Itoa(j))
		next[i] = j
	}
}
func StrStr1(content string, model string) int {
	if len(model) == 0 {
		return 0
	}
	//  构建前缀表
	next := make([]int, len(model))
	getNext1(next, model)
	j := -1 // 模式串的起始位置 next为-1 因此也为-1
	for i := 0; i < len(content); i++ {
		for j >= 0 && content[i] != model[j+1] {
			j = next[j] // 寻找下一个匹配点
		}
		if content[i] == model[j+1] {
			j++
		}
		if j == len(model)-1 { // j指向了模式串的末尾
			return i - len(model) + 1
		}
	}
	return -1
}

// 方法二: 前缀表无减一或者右移

// getNext2 构造前缀表next
// params:
//
//	next 前缀表数组
//	s 模式串
func getNext2(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
func StrStr2(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext2(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}
