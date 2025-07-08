package main

func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}

	prefix := str[0]
	shortestLength := len(prefix)
	for _, s := range str {
		if len(s) < shortestLength {
			shortestLength = len(s)
		}
	}

	for i := 1; i < shortestLength; i++ {
		for j := 1; j < len(str); j++ {
			if str[j][i] != prefix[i] {
				return prefix[:i]
			}
		}
	}

	return prefix
}

func main() {
	commonStrings := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(commonStrings)
	println("Longest common prefix:", result) // Output: "fl"
}
