package backpack

import (
	"sort"
	"strings"
)

//如果一个密码满足以下所有条件，我们称它是一个 强 密码：
//
//它有至少 8 个字符。
//至少包含 一个小写英文 字母。
//至少包含 一个大写英文 字母。
//至少包含 一个数字 。
//至少包含 一个特殊字符 。特殊字符为："!@#$%^&*()-+" 中的一个。
//它 不 包含 2 个连续相同的字符（比方说 "aab" 不符合该条件，但是 "aba" 符合该条件）。
//给你一个字符串 password ，如果它是一个 强 密码，返回 true，否则返回 false 。

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	for i := 1; i < len(password); i++ {
		if password[i] == password[i-1] {
			return false
		}
	}
	var (
		a, b, c, d int
	)
	for _, v := range password {
		if v >= 'a' && v <= 'z' {
			a++
		}
		if v >= 'A' && v <= 'Z' {
			b++
		}
		if v > '0' && v < '9' {
			c++
		}
		if strings.Contains("!@#$%^&*()-+", string(v)) {
			d++
		}
	}

	if a < 1 || b < 1 || c < 1 || d < 1 {
		return false
	}
	return true
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	var ans []int = make([]int, len(spells))
	sort.Ints(potions)
	for i, v := range spells {
		left, right := 0, len(potions)-1
		var mid int
		for left < right {
			mid = (left+right)/2 + 1
			tt := int64(potions[mid] * v)
			if tt < success {
				left = mid + 1
				continue
			}
			right = mid - 1

		}

		ans[i] = len(potions) - left - 1
	}
	return ans
}

func calculateTax(brackets [][]int, income int) float64 {
	if len(brackets) < 1 {
		return 0
	}
	if len(brackets) == 1 {

		if income > brackets[0][0] {
			return float64(brackets[0][0]*brackets[0][1]) / 100
		}

		return float64(income*brackets[0][1]) / 100
	}

	// 小于第一个
	if income < brackets[0][0] {
		return float64(income*brackets[0][1]) / 100
	}

	// 大于第一个
	total := brackets[0][1] * brackets[0][0]
	sum := brackets[0][0]
	for i := 1; i < len(brackets); i++ {
		if income < brackets[i][0] {
			total += (income - sum) * brackets[i][1]
			break
		}
		total += (brackets[i][0] - brackets[i-1][0]) * brackets[i][1]
		sum = brackets[i][0]
	}
	return float64(total) / 100

}
