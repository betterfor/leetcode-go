package code

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}

	var dp = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m+1; i++ {
		dp[i][0] = i
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown += 1
			}
			var min = left
			if min > down {
				min = down
			}
			if min > leftDown {
				min = leftDown
			}
			dp[i][j] = min
		}
	}
	return dp[m][n]
}
