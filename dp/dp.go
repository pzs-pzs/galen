package dp

import "strings"

func longestCommonSubsequence(s string, w string) int {
	m, n := len(s), len(w)
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == w[j] {
				if i-1 >= 0 && j-1 >= 0 {
					ans[i][j] = ans[i-1][j-1] + 1
					continue
				}
				ans[i][j] = 1
				continue
			}
			if j-1 < 0 && i-1 >= 0 {
				ans[i][j] = ans[i-1][j]
				continue
			}
			if j-1 >= 0 && i-1 < 0 {
				ans[i][j] = ans[i][j-1]
				continue
			}
			if i-1 < 0 && j-1 < 0 {
				ans[i][j] = 0
				continue
			}
			ans[i][j] = max(ans[i-1][j], ans[i][j-1])

		}
	}

	return ans[m-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonPrefix(strs []string) string {
	var ans strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		flag := true
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[0][i] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		ans.WriteRune(rune(strs[0][i]))
	}
	return ans.String()
}
