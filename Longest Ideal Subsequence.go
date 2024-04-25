func solve(i, prev, k int, s string, dp [][]int) int {
    if i == len(s) {
        return 0
    }

    if dp[i][prev+1] != -1 {
        return dp[i][prev+1]
    }

    notTake := 0 + solve(i+1, prev, k, s, dp)
    Take := 0
    if prev == -1 || int(math.Abs(float64(s[i]-'a')-float64(prev))) <= k {
        Take = 1 + solve(i+1, int(s[i]-'a'), k, s, dp)
    }

    dp[i][prev+1] = max(notTake, Take)
    return dp[i][prev+1]
}

func longestIdealString(s string, k int) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, 28)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    return solve(0, -1, k, s, dp)
}
