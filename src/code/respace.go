package main

func respace(dictionary []string, sentence string) int {
	dictMap := make(map[string]int)

	for _, s := range dictionary {
		dictMap[s] = 1
	}

	dp := make([]int, len(sentence)+1)

	for i := 1; i < len(sentence); i++ {
		// 考虑如果不可以匹配 + 1很关键
		dp[i] = dp[i-1] + 1

		for j := 0; j < i; j++ {
			if dictMap[sentence[j:i]] == 1 {
				if dp[j] < dp[i] {
					dp[i] = dp[j]
				}
			}
		}
	}
	return dp[len(sentence)]
}
