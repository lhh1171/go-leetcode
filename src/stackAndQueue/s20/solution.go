package s20

func IsValid(s string) bool {
	//  1.创建map
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	//  2.利用栈把前括号装进去，遇见后括号就检测栈看看栈顶是不是前括号
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
