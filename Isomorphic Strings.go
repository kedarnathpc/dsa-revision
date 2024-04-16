func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sv := make([]int, 128)
    tv := make([]int, 128)

    for i := range sv {
        sv[i] = -1
        tv[i] = -1
    }

    for i := 0; i < len(s); i++ {
        if sv[s[i]] != tv[t[i]] {
            return false
        }
        sv[s[i]] = i
        tv[t[i]] = i
    }

    return true
}
