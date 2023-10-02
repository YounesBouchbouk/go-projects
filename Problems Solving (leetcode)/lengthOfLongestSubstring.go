package main

// https://www.code-recipe.com/post/longest-substring-without-repeating-characters

func lengthOfLongestSubstring(s string) int {

	start := 0

	var charIndexMap = make(map[byte]int)

	result := 0

	for end := 0; end < len(s); end++ {

		duplicateIndex, isDuplicate := charIndexMap[s[end]]

		if isDuplicate {

			result = max(result, end-start)

			for i := start; i <= duplicateIndex; i++ {
				delete(charIndexMap, s[i])
			}

			start = duplicateIndex + 1

		}

		charIndexMap[s[end]] = end

	}
	result = max(result, len(s)-start)

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
