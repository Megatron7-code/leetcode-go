package main

/**
	link: https://leetcode-cn.com/problems/reverse-integer/
	desc: 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 */
const INT32_MAX = int(^uint32(0) >> 1)
const INT32_MIN = ^INT32_MAX

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		rev = rev * 10 + pop
		if (rev > INT32_MAX || (rev == INT32_MAX/10 && pop > INT32_MAX%10)) {
			return 0
		}
		if (rev < INT32_MIN || (rev == INT32_MIN/10 && pop < INT32_MIN%10)) {
			return 0
		}
	}
	return rev
}

func main(){
	a := 2147483648
	println(reverse(a))
}