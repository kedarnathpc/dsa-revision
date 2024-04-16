func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)
    n := len(strs)
    res := ""
    first, last := strs[0], strs[n-1]

    for i := 0; i < len(first); i++ {
        if first[i] == last[i] {
            res += string(first[i])
        } else {
            break
        }
    }

    return res
}
