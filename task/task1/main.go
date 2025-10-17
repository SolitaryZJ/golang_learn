package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(oneTimeNum([]int{1, 2, 1, 3, 3}))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isValid("(){}[]"))
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(plusOne([]int{1, 2, 9}))
	fmt.Println(removeDuplicates([]int{1, 1, 2, 2, 3}))
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

// 两数之和
func twoSum(nums []int, target int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		sub := target - v
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == sub {
				ans = append(ans, i, j)
				return ans
			}
		}
	}
	return ans
}

// 合并区间
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	for _, interval := range intervals {
		m := len(ans)
		if m > 0 && interval[0] <= ans[m-1][1] {
			ans[m-1][1] = max(ans[m-1][1], interval[1])
		} else {
			ans = append(ans, interval)
		}
	}
	return ans
}

// 移除数组中的重复元素
// 思路：元素作为map key，如果map中已经存在，则返回
func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k

	//return len(slices.Compact(nums))
}

// 大数+1.
// 思路：从右到左找到第一个非9的数字，然后+1，如果所有数字是9，则将所有数字都置为0，然后+1
func plusOne(digits []int) []int {
	for i, d := range slices.Backward(digits) {
		if d < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	digits = append(digits, 0)
	digits[0] = 1
	return digits
}

// 找出最长的公共前缀
// 思路：从左到右，从上到下遍历
func longestCommonPrefix(strs []string) string {
	str0 := strs[0]
	// 从左到右
	for j, v := range str0 {
		// 从上到下
		for _, str := range strs {
			if len(str) == j || str[j] != byte(v) {
				return str0[:j]
			}
		}
	}
	return str0
}

func isValid(s string) bool {
	// 必须偶数长度
	if len(s)%2 != 0 {
		return false
	}
	mir := map[rune]rune{')': '(', ']': '[', '}': '{'}
	// 使用切片作为栈
	stack := []rune{}
	for _, v := range s {
		// 左括号入栈
		if mir[v] == 0 {
			stack = append(stack, v)
		} else { // 右括号出栈
			if len(stack) == 0 || stack[len(stack)-1] != mir[v] {
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		}
	}
	return len(stack) == 0
}

func isPalindrome(x int) bool {
	if x < 0 || x > 0 && x%10 == 0 {
		return false
	}
	reverse := 0
	for reverse < x/10 {
		reverse = reverse*10 + x%10
		x /= 10
	}
	return reverse == x || reverse == x/10
}

func oneTimeNum(arr []int) int {
	var times = make(map[int](int))
	for _, v := range arr {
		times[v] = times[v] + 1
	}

	for k, v := range times {
		if v == 1 {
			return k
		}
	}
	return 0
}
