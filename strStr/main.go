package main

/**
	link: https://leetcode-cn.com/problems/implement-strstr/
	dest: 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
 */

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	needleLen := len(needle)
	haystackLen := len(haystack)
	if needleLen > haystackLen {
		return -1
	}
	for i := 0; i < haystackLen; i++ {
		if i + needleLen > haystackLen {
			return -1
		}
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

func main(){
	haystack := "hello, world"
	needle := "ll"
	print(strStr(haystack, needle))
}
