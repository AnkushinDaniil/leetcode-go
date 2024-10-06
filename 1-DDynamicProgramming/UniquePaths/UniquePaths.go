package uniquepaths

func uniquePaths(m int, n int) int {
    ans := 1
    m--
    n--
    for i := 1; i <= m ; i++ {
        ans = ans * (n + i) / i
    }
    return ans
}
