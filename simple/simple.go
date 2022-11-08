package simple

func countConsistentStrings(allowed string, words []string) int {
	cache := map[rune]struct{}{}
	for _, v := range allowed {
		cache[v] = struct{}{}
	}

	var ans int
	for _, word := range words {
		flag := true
		for _, v := range word {
			_, ok := cache[v]
			if !ok {
				flag = false
			}
		}
		if flag {
			ans++
		}
	}

	return ans
}
